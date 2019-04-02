package eventdata

import (
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestGetEventData(t *testing.T) {
	type args struct {
		jsonBytes []byte
	}
	tests := []struct {
		name     string
		args     args
		wantData EventData
		wantErr  bool
		loadErr  bool
	}{
		{
			name: "test1",
			args: args{
				jsonBytes: []byte{},
			},
			wantData: EventData{
				ID:    5212,
				URL:   `https://wp.infra-workshop.tech?p=5212`,
				Title: "\u30a4\u30f3\u30d5\u30e9\u30a8\u30f3\u30b8\u30cb\u30a2\u306e\u305f\u3081\u306e\u30c0\u30a4\u30a8\u30c3\u30c8",
				Description: `1.概要
`,
				StartDate: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				Organizer: "",
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				jsonBytes: []byte{},
			},
			wantData: EventData{
				ID:    5212,
				URL:   `https://wp.infra-workshop.tech?p=5212`,
				Title: "\u30a4\u30f3\u30d5\u30e9\u30a8\u30f3\u30b8\u30cb\u30a2\u306e\u305f\u3081\u306e\u30c0\u30a4\u30a8\u30c3\u30c8",
				Description: `1.概要
`,
				StartDate: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				Organizer: "",
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	// loading json-data
	for n := range tests {
		raw, err := ioutil.ReadFile("EventData_testcase/" + tests[n].name + ".json")
		if err != nil {
			t.Errorf("Can't load json : %v / err : %v", tests[n].name, err)
			tests[n].loadErr = true
		} else {
			tests[n].loadErr = false
		}
		tests[n].args.jsonBytes = raw
	}
	// test execute
	for _, tt := range tests {
		if tt.loadErr {
			t.Logf("%v is Skipping...", tt.name)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := GetEventData(tt.args.jsonBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetEventData() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}

func TestGetEventsFromWordpress(t *testing.T) {
	type args struct {
		url         string
		dayLineHour int
	}
	tests := []struct {
		name       string
		args       args
		wantEvents []EventData
		wantErr    bool
	}{
		{
			name: "test1",
			args: args{
				"wp.infra-workshop.tech", 4,
			},
			wantEvents: []EventData{{
				ID:          5212,
				URL:         "wp.infra-workshop.tech?id=5212",
				Title:       "\u30a4\u30f3\u30d5\u30e9\u30a8\u30f3\u30b8\u30cb\u30a2\u306e\u305f\u3081\u306e\u30c0\u30a4\u30a8\u30c3\u30c8",
				Description: "",
			}},
			wantErr: false,
		},
		{
			name: "test2 cache-test",
			args: args{
				"wp.infra-workshop.tech", 4,
			},
			wantEvents: []EventData{{
				ID:          5212,
				URL:         "wp.infra-workshop.tech?id=5212",
				Title:       "\u30a4\u30f3\u30d5\u30e9\u30a8\u30f3\u30b8\u30cb\u30a2\u306e\u305f\u3081\u306e\u30c0\u30a4\u30a8\u30c3\u30c8",
				Description: "",
			}},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetEventsFromWordpress(tt.args.url, tt.args.dayLineHour)
			// gotEvents, err := GetEventsFromWordpress(tt.args.url, tt.args.dayLineHour)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventsFromWordpress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// if !reflect.DeepEqual(gotEvents, tt.wantEvents) {
			// 	t.Errorf("GetEventsFromWordpress() = %v, want %v", gotEvents, tt.wantEvents)
			// }
		})
	}
}