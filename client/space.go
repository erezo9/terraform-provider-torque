package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (c *Client) CreateSpace(name string, color string, icon string) error {
	fmt.Println(c.HostURL + "api/spaces")

	space := Space{
		Name:  name,
		Color: color,
		Icon:  icon,
	}

	payload, err := json.Marshal(space)
	if err != nil {
		log.Fatalf("impossible to marshall space: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/spaces", c.HostURL), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteSpace(name string) error {
	fmt.Println(c.HostURL + "api/spaces")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%sapi/spaces/%s", c.HostURL, name), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddAgentToSpace(agent string, ns string, sa string, space string, agnet_type string) error {
	fmt.Println(c.HostURL + "api/spaces")

	data := AgentSpaceAssociation{
		Type:                  agnet_type,
		DefaultNamespace:      ns,
		DefaultServiceAccount: sa,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall agent association: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/spaces/%s/agents/%s", c.HostURL, space, agent), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveAgentFromSpace(agent string, space string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%sapi/spaces/%s/agents/%s", c.HostURL, space, agent), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) OnboardRepoToSpace(space_name string, repo_name string, repo_type string, repo_url string, repo_token string, repo_branch string) error {
	fmt.Println(c.HostURL + "api/spaces")

	data := RepoSpaceAssociation{
		Type:        repo_type,
		URL:         repo_url,
		AccessToken: repo_token,
		Branch:      repo_branch,
		Name:        repo_name,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall agent association: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/spaces/%s/repositories", c.HostURL, space_name), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveRepoFromSpace(space_name string, repo_name string) error {
	fmt.Println(c.HostURL + "api/spaces")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%sapi/spaces/%s/repositories?repository_name=%s", c.HostURL, space_name, repo_name), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

// /api/spaces/{space_name}/blueprints.
func (c *Client) GetSpaceBlueprints(space_name string) ([]Blueprint, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%sapi/spaces/%s/blueprints", c.HostURL, space_name), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return nil, err
	}

	blueprints := []Blueprint{}
	err = json.Unmarshal(body, &blueprints)
	if err != nil {
		return nil, err
	}

	fmt.Println("Blueprint list length: " + strconv.Itoa(len(blueprints)))

	return blueprints, nil
}

func (c *Client) SetSpaceTagValue(space_name string, tag_name string, tag_value string) error {
	fmt.Println(c.HostURL + "api/spaces")

	data := TagNameValue{
		Name:  tag_name,
		Value: tag_value,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall space tag key value association: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/spaces/%s/settings/tags", c.HostURL, space_name), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SetBlueprintTagValue(space_name string, tag_name string, tag_value string, repo_name string, blueprint_name string) error {
	fmt.Println(c.HostURL + "api/spaces")

	data := TagNameValue{
		Name:  tag_name,
		Value: tag_value,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall blueprint tag key value association: %s", err)
	}
	// /api/spaces/devnet/repositories/qtorque/blueprints/Elasticsearch/settings/tags
	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/spaces/%s/repositories/%s/blueprints/%s/settings/tags", c.HostURL, space_name, repo_name, blueprint_name), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) PublishBlueprintInSpace(space_name string, repo_name string, blueprint_name string) error {
	data := CatalogItemRequest{
		BlueprintName:  blueprint_name,
		RepositoryName: repo_name,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall agent association: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/spaces/%s/catalog", c.HostURL, space_name), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UnpublishBlueprintInSpace(space_name string, repo_name string, blueprint_name string) error {
	fmt.Println(c.HostURL + "api/spaces")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%sapi/spaces/%s/catalog/%s?repository_name=%s", c.HostURL, space_name, blueprint_name, repo_name), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddSpaceParameter(space_name string, name string, value string, sensitive bool, description string) error {
	data := ParameterRequest{
		Name:        name,
		Value:       value,
		Sensitive:   sensitive,
		Description: description,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall agent association: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/spaces/%s/settings/parameters", c.HostURL, space_name), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteSpaceParameter(space_name string, parameter_name string) error {
	fmt.Println(c.HostURL + "api/spaces")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%sapi/spaces/%s/settings/parameters/%s", c.HostURL, space_name, parameter_name), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddAccountParameter(name string, value string, sensitive bool, description string) error {
	data := ParameterRequest{
		Name:        name,
		Value:       value,
		Sensitive:   sensitive,
		Description: description,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall agent association: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/settings/parameters", c.HostURL), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetSpaceParameter(space_name string, parameter_name string) (ParameterRequest, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%sapi/spaces/%s/settings/parameters", c.HostURL, space_name), nil)
	if err != nil {
		return ParameterRequest{}, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return ParameterRequest{}, err
	}

	params := []ParameterRequest{}
	err = json.Unmarshal(body, &params)
	if err != nil {
		return ParameterRequest{}, err
	}

	param := ParameterRequest{}
	for _, n := range params {
		if parameter_name == n.Name {
			param = n
		}
	}

	return param, nil
}

func (c *Client) GetAccountParameter(parameter_name string) (ParameterRequest, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%sapi/settings/parameters", c.HostURL), nil)
	if err != nil {
		return ParameterRequest{}, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return ParameterRequest{}, err
	}

	params := []ParameterRequest{}
	err = json.Unmarshal(body, &params)
	if err != nil {
		return ParameterRequest{}, err
	}

	param := ParameterRequest{}
	for _, n := range params {
		if parameter_name == n.Name {
			param = n
		}
	}

	return param, nil
}

func (c *Client) DeleteAccountParameter(parameter_name string) error {
	fmt.Println(c.HostURL + "api/spaces")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%sapi/settings/parameters/%s", c.HostURL, parameter_name), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddGroupToSpace(groupName string, description string, idpId string, users []string, accountRole string, spaceRole []SpaceRole) error {

	data := GroupRequest{
		Name:        groupName,
		Description: description,
		IdpId:       idpId,
		Users:       users,
		AccountRole: accountRole,
		SpaceRoles:  spaceRole,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall group request: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/groups", c.HostURL), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteGroup(group_name string) error {
	fmt.Println(c.HostURL + "api/spaces")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%sapi/groups/%s", c.HostURL, group_name), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddAWSCostTarget(name string, target_type string, role_arn string, external_id string) error {

	data := AwsCostTaret{
		Name:       name,
		Type:       target_type,
		ARN:        role_arn,
		ExternalId: external_id,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall aws cost target request: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%sapi/settings/costtargets", c.HostURL), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteCostTarget(target_name string) error {
	fmt.Println(c.HostURL + "api/spaces")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%sapi/settings/costtargets/%s", c.HostURL, target_name), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetGroup(group_name string) (GroupRequest, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%sapi/groups", c.HostURL), nil)
	if err != nil {
		return GroupRequest{}, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return GroupRequest{}, err
	}

	groups := []GroupRequest{}
	err = json.Unmarshal(body, &groups)
	if err != nil {
		return GroupRequest{}, err
	}

	group := GroupRequest{}
	for _, n := range groups {
		if group_name == n.Name {
			group = n
		}
	}

	return group, nil
}

func (c *Client) GetSpace(space_name string) (Space, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%sapi/spaces/%s", c.HostURL, space_name), nil)
	if err != nil {
		return Space{}, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return Space{}, err
	}

	space := Space{}
	err = json.Unmarshal(body, &space)
	if err != nil {
		return Space{}, err
	}

	return space, nil
}

func (c *Client) GetSpaces() ([]Space, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%sapi/spaces", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token)
	if err != nil {
		return nil, err
	}

	spaces := []Space{}
	err = json.Unmarshal(body, &spaces)
	if err != nil {
		return nil, err
	}

	return spaces, nil
}

func (c *Client) UpdateAccountTag(name string, value string, description string, possible_values []string, scope string) error {

	tag := Tag{
		Name:           name,
		Value:          value,
		Scope:          scope,
		Description:    description,
		PossibleValues: possible_values,
	}

	payload, err := json.Marshal(tag)
	if err != nil {
		log.Fatalf("impossible to marshall update group request: %s", err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%sapi/settings/tags/%s", c.HostURL, name), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateAccountParameter(name string, value string, sensitive bool, description string) error {

	data := ParameterRequest{
		Name:        name,
		Value:       value,
		Sensitive:   sensitive,
		Description: description,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall update group request: %s", err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%sapi/settings/parameters/%s", c.HostURL, name), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateGroup(groupName string, description string, idpId string, users []string, accountRole string, spaceRole []SpaceRole) error {

	data := GroupRequest{
		Name:        groupName,
		Description: description,
		IdpId:       idpId,
		Users:       users,
		AccountRole: accountRole,
		SpaceRoles:  spaceRole,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall update group request: %s", err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%sapi/groups/%s", c.HostURL, groupName), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

//  https://portal.qtorque.io/api/spaces/events-ms
// {"name":"events-ms","users_count":34,"groups_count":1,"color":"darkBlue","icon":"cloud","spaceName":"events-ms"}
// {
// 	"name": "10-CloudShell",
// 	"users_count": 34,
// 	"groups_count": 1,
// 	"color": "darkBlue",
// 	"icon": "screen"
// },

func (c *Client) UpdateSpace(name string, color string, icon string) error {

	data := Space{
		Name:  name,
		Color: color,
		Icon:  icon,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("impossible to marshall update space request: %s", err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%sapi/spaces/%s", c.HostURL, name), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}
