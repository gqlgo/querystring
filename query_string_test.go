package main

import (
	"bytes"
	_ "embed"
	"os"
	"testing"
)

func TestQueryString(t *testing.T) {
	type args struct {
		fileName string
	}
	type want struct {
		queryFileName string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "only query",
			args: args{
				fileName: "testdata/test1.tsx",
			},
			want: want{
				queryFileName: "testdata/test1.graphql",
			},
		},
		{
			name: "query and mutation",
			args: args{
				fileName: "testdata/test2.tsx",
			},
			want: want{
				queryFileName: "testdata/test2.graphql",
			},
		},
		{
			name: "query, mutation, and other text",
			args: args{
				fileName: "testdata/test3.tsx",
			},
			want: want{
				queryFileName: "testdata/test3.graphql",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argB, err := os.ReadFile(tt.args.fileName)
			if err != nil {
				t.Errorf("Read %v failed", tt.args.fileName)
			}

			// Prepare want format query
			wantB, err := os.ReadFile(tt.want.queryFileName)
			if err != nil {
				t.Errorf("Read %v failed", tt.want.queryFileName)
			}
			wantFormatQuery, _ := formatQuery(wantB)

			// got format query
			got := queryString(argB)
			if !bytes.Equal(wantFormatQuery, got) {
				t.Errorf("want:\n%s\n got:\n%s", wantFormatQuery, got)
			}
		})
	}
}
