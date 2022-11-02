package aftership

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCanonicalizedResource(t *testing.T) {
	type args struct {
		rawUrl string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{
			name: "/tracking",
			args: args{
				rawUrl: "https://api.aftership.com/tracking",
			},
			wantResult: "/tracking",
			wantErr:    false,
		},
		{
			name: "/tracking",
			args: args{
				rawUrl: "https://api.aftership.com/tracking?3key=111&2key=222&key=333&1key=000",
			},
			wantResult: "/tracking?1key=000&2key=222&3key=111&key=333",
			wantErr:    false,
		},
	}
	for _, cur := range tests {
		tt := cur
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetCanonicalizedResource(tt.args.rawUrl)
			assert.Equal(t, tt.wantResult, gotResult)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGetCanonicalizedHeaders(t *testing.T) {
	type args struct {
		headers map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetCanonicalizedHeaders",
			args: args{
				headers: map[string]string{
					"AS-header2": " ThisIsHeader2",
					"AS-Header1": " this-is-header-1",
				},
			},
			want: "as-header1:this-is-header-1\nas-header2:ThisIsHeader2",
		},
	}
	for _, cur := range tests {
		tt := cur
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetCanonicalizedHeaders(tt.args.headers))
		})
	}
}

func TestGetSignString(t *testing.T) {
	type args struct {
		method                 string
		body                   string
		contentType            string
		date                   string
		canonicalizedAmHeaders string
		canonicalizedResource  string
	}
	dateStr := time.Now().UTC().Format(http.TimeFormat)
	tests := []struct {
		name          string
		args          args
		wantSignature string
		wantErr       bool
	}{
		{
			name: "GET",
			args: args{
				method:                 "GET",
				body:                   "",
				contentType:            "application/json",
				date:                   dateStr,
				canonicalizedAmHeaders: "as-header1:this-is-header-1\nas-header2:ThisIsHeader2",
				canonicalizedResource:  "/tracking?1key=000&2key=222&3key=111&key=333",
			},
			wantSignature: "GET\n\n\n" + dateStr + "\nas-header1:this-is-header-1\nas-header2:ThisIsHeader2\n/tracking?1key=000\u00262key=222\u00263key=111\u0026key=333",
			wantErr:       false,
		},
		{
			name: "POST",
			args: args{
				method:                 "POST",
				body:                   `{"key1":"value1"}`,
				contentType:            "application/json",
				date:                   dateStr,
				canonicalizedAmHeaders: "as-header1:this-is-header-1\nas-header2:ThisIsHeader2",
				canonicalizedResource:  "/tracking?1key=000&2key=222&3key=111&key=333",
			},
			wantSignature: "POST\n1530ACCF30CAB891D759FA3CB8322250\napplication/json\n" + dateStr + "\nas-header1:this-is-header-1\nas-header2:ThisIsHeader2\n/tracking?1key=000\u00262key=222\u00263key=111\u0026key=333",
			wantErr:       false,
		},
	}
	for _, cur := range tests {
		tt := cur
		t.Run(tt.name, func(t *testing.T) {
			gotSignature, err := GetSignString(tt.args.method, tt.args.body, tt.args.contentType, tt.args.date, tt.args.canonicalizedAmHeaders, tt.args.canonicalizedResource)
			assert.Equal(t, tt.wantSignature, gotSignature)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
