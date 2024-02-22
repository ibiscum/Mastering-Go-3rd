package main

import (
	"reflect"
	"testing"
)

func Test_readCSVFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCSVFile(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCSVFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCSVFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saveCSVFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveCSVFile(tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("saveCSVFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
