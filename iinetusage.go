package iinetusage

import (
  "encoding/xml"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"
)

var netClient = &http.Client{
  Timeout: time.Second * 20,
}

type result struct {
  XMLName      xml.Name      `xml:"ii_feed"`
  Quotas       []quotaReset  `xml:"volume_usage>quota_reset"`
  TrafficTypes []trafficType `xml:"volume_usage>expected_traffic_types>type"`
}

type quotaReset struct {
  Anniversary   uint64 `xml:"anniversary"`
  DaysRemaining uint64 `xml:"days_remaining"`
}

type trafficType struct {
  Classification string  `xml:"classification,attr"`
  Used           uint64  `xml:"used,attr"`
  Quotas         []quota `xml:"quota_allocation"`
}

type quota struct {
  Amount uint64 `xml:",chardata"`
}

// IINetUsage ...
type IINetUsage struct {
  Quota            uint64  `json:"quota"`
  Used             uint64  `json:"used"`
  Remaining        uint64  `json:"remaining"`
  PercentUsed      float64 `json:"percent_used"`
  PercentRemaining float64 `json:"percent_remaining"`
  DaysRemaining    uint64  `json:"days_remaining"`
}

// IINet ...
type IINet struct {
  username string
  password string
}

// New ...
func New(username string, password string) *IINet {
  iinet := &IINet{
    username: username,
    password: password,
  }

  return iinet
}

// GetUsage ...
func (iinet *IINet) GetUsage() (*IINetUsage, error) {
  var usage *IINetUsage

  r, err := getResult(iinet.username, iinet.password)
  if err != nil {
    return usage, err
  }

  if len(r.Quotas) == 0 {
    return usage, errors.New("Could not parse XML, check your credentials")
  }

  daysRemaining := r.Quotas[0].DaysRemaining
  quota := r.TrafficTypes[0].Quotas[0].Amount * 1000000
  used := r.TrafficTypes[0].Used
  percentUsed := (float64(used) / float64(quota)) * 100
  percentRemaining := 100 - percentUsed

  usage = &IINetUsage{
    Quota:            quota,
    Used:             used,
    Remaining:        quota - used,
    PercentUsed:      percentUsed,
    PercentRemaining: percentRemaining,
    DaysRemaining:    daysRemaining,
  }

  return usage, nil
}

func getResult(username string, password string) (result, error) {
  r := result{}
  url := fmt.Sprintf("https://toolbox.iinet.net.au/cgi-bin/new/volume_usage_xml.cgi?username=%s&action=login&password=%s", username, password)

  response, err := netClient.Get(url)
  if err != nil {
    return r, err
  }
  data, _ := ioutil.ReadAll(response.Body)
  response.Body.Close()

  err = xml.Unmarshal([]byte(data), &r)
  if err != nil {
    return r, err
  }

  return r, nil
}
