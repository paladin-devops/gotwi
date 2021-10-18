package types_test

import (
	"testing"

	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetTimelinesTweetsParams_SetAccessToken(t *testing.T) {
	cases := []struct {
		name   string
		token  string
		expect string
	}{
		{
			name:   "normal",
			token:  "test-token",
			expect: "test-token",
		},
		{
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.TweetTimelinesTweetsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_TweetTimelinesTweetsParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.TweetTimelinesTweetsParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.TweetTimelinesTweetsParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "with exclude",
			params: &types.TweetTimelinesTweetsParams{
				ID:      "test-id",
				Exclude: []string{"exc1", "exc2"},
			},
			expect: endpointRoot + "test-id" + "?exclude=exc1%2Cexc2",
		},
		{
			name: "with expansions",
			params: &types.TweetTimelinesTweetsParams{
				ID:         "test-id",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with max_results",
			params: &types.TweetTimelinesTweetsParams{
				ID:         "test-id",
				MaxResults: 50,
			},
			expect: endpointRoot + "test-id" + "?max_results=50",
		},
		{
			name: "with media.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:          "test-id",
				MediaFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:          "test-id",
				PlaceFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:         "test-id",
				PollFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:          "test-id",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:         "test-id",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.TweetTimelinesTweetsParams{
				ID:          "test-id",
				Exclude:     []string{"exc"},
				Expansions:  []string{"ex"},
				MaxResults:  50,
				MediaFields: []string{"mf"},
				PlaceFields: []string{"plf"},
				PollFields:  []string{"pof"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointRoot + "test-id" + "?exclude=exc&expansions=ex&max_results=50&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.TweetTimelinesTweetsParams{
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_TweetTimelinesTweetsParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetTimelinesTweetsParams
	}{
		{
			name:   "empty params",
			params: &types.TweetTimelinesTweetsParams{},
		},
		{
			name:   "some params",
			params: &types.TweetTimelinesTweetsParams{ID: "id"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}