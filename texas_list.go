package main

import (
  "flag"
   "github.com/golang/glog"
   "github.com/openedinc/opened-go"

)

// init parses the flag options
func init() {
  flag.Parse()
}
func main() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://partner.opened.com/1/standard_groups.json", nil)

	token,err := opened.GetToken ("","","","")

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
}
func main() {
  glog.Infoln("Texas Standards")
token,err := opened.GetToken("","","","")
  query_params:=make(map[string]string)
  query_params["descriptive"]="fractions"
  query_params["standard_group"]="22"
  query_params["grades_range"]="K-12"
  results,err:=opened.SearchResources(query_params,token)
  if err!=nil {
    glog.Errorf("Error from SearchResources %+v",err)
  }
  glog.V(1).Infof("%d results returned",len(results.Resources))
  for _,resource:= range results.Resources {
    //I want the output to simply be "standard id, count number of resources"
    glog.V(2).Infof("Resource %s (%d)",resource.Id,resource.Title)
  }
}
