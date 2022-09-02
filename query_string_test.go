package main

import (
	"bytes"
	_ "embed"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		fileName string
	}
	type want struct {
		queryFileName string
		err           error
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
			wantFormatQuery, err := formatQuery(wantB)
			if err != nil {
				t.Errorf("want format failed: %v\n%s", err, wantB)
			}

			// got format query
			gotFormatQuery, err := queryString(argB)
			if err != tt.want.err {
				t.Errorf("want:%v got:%v", tt.want.err, err)
			}
			if !bytes.Equal(wantFormatQuery, gotFormatQuery) {
				t.Errorf("want:\n%s\n got:\n%s", wantFormatQuery, gotFormatQuery)
			}
		})
	}
}
