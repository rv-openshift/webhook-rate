// not yet finish...

package main

import (
	"net/http/httptest"
)

type m_struct struct {                    // json for gitlab api payload                         
	Id              string        `json:"id"`
	Short_id        string        `json:"short_id"`
	Created_at      string        `json:"created_at"`
	Parent_ids      []interface{} `json:"parent_ids"`
	Title           string        `json:"title"`
	Message         string        `json:"message"`
	Author_name     string        `json:"author_name"`
	Author_email    string        `json:"author_email"`
	Authored_date   string        `json:"authored_date"`
	Committer_name  string        `json:"committer_name"`
	Committer_email string        `json:"committer_email"`
	Committed_date  string        `json:"committed_date"`
	Trailers        struct {
		Primary string `json:"primary"`
	} `json:"unknown,omitempty"`
	Web_url       string `json:"web_url"`
	Status        string `json:"status"`
	Project_id    int    `json:"project_id"`
	Last_pipeline struct {
		Id         int    `json:"id"`
		Iid        int    `json:"iid"`
		Project_id int    `json:"project_id"`
		Sha        string `json:"sha"`
		Ref        string `json:"ref"`
		Status     string `json:"status"`
		Source     string `json:"source"`
		Created_at string `json:"created_at"`
		Updated_at string `json:"Updated"`
		Web_url    string `json:"web_url"`
	} `json:"last_pipeline"`
	Stats struct {
		Additions int `json:"additions"`
		Deletions int `json:"deletions"`
		Total     int `json:"total"`
	} `json:"stats"`
}

type Tests struct {
	name string
	server *httptest.server
	response *m_struct
	expcetedError error
}

func TesthttpGet (t *testing.T) {
	test := []Tests {
		name: "gilatb-api"
		server: httptest.NewMain(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			w.WriteHeader(http.StatusOK)	
			w.Write([]byte(`{??????????????????}`)),
			response: &m_struct{
				Id: "46c3e11e77912f44746f7727047cbf2176d3cf14",
				Short_id: "46c3e11e",
				Created_at: "2023-10-16T12:31:00.000-05:00",
				Parent_ids: [
					"a08279e15b87d07c2ca288180824a262bd20608c",
					"09667205a6f4c54178ae9a97a2cc89886cc20cb7"
					],
				Title: "gitops: qa-1.3.9 by ricarte.venerayan@toronto.ca",
				Message": "gitops: qa-1.3.9 by ricarte.venerayan@toronto.ca",
				Author_name: "cd-token",
				Author_email: "group_10_bot_7dc56aeb109e98f40798ab6fce76c4e1@noreply.cd-sbx.toronto.ca",
				Authored_date: "2023-10-16T12:31:00.000-05:00",
				Committer_name: "cd-token",
				Committer_email: "group_10_bot_7dc56aeb109e98f40798ab6fce76c4e1@noreply.cd-sbx.toronto.ca",
				Committed_date: "2023-10-16T12:31:00.000-05:00",
				Trailers: {},
					Web_url: "https://cd-sbx.toronto.ca:6443/devops/eis/tsd-data/-/commit/46c3e11e77912f44746f7727047cbf2176d3cf14",
					Stats: {
					Additions: 2,
					Deletions: 2,
					Total": 4
					},
				Status: null,
				Project_id: 21,
				Lst_pipeline: null
				},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		t.RUn(test.name, func(t *testing.T){
		// need to complete this
		}
	}
}