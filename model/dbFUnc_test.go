package model

import (
	"reflect"
	"testing"
)

func TestConfirmRemindExist(t *testing.T) {
	type args struct {
		uid string
	}
	Db.Init()
	defer Db.Close()
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				uid: "2019214300",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfirmRemindExist(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfirmRemindExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConfirmRemindExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfirmRequirementExist(t *testing.T) {
	Db.Init()
	defer Db.Close()
	type args struct {
		requirements Requirements
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 bool
	}{
		{
			name: "1",
			args: args{
				requirements: Requirements{
					RequirementId:    0,
					SenderSid:        "2019214300",
					Title:            "first",
					Content:          "lalalalalala",
					PostTime:         "",
					Date:             0,
					TimeFrom:         0,
					TimeEnd:          0,
					RequirePeopleNum: 0,
					Place:            0,
					Tag:              0,
					Type:             0,
					ContactWayType:   "",
					ContactWay:       "",
					Status:           0,
				},
			},
			want:  nil,
			want1: true,
		},
		{
			name: "2",
			args: args{
				requirements: Requirements{
					RequirementId:    0,
					SenderSid:        "2019",
					Title:            "",
					Content:          "",
					PostTime:         "",
					Date:             0,
					TimeFrom:         0,
					TimeEnd:          0,
					RequirePeopleNum: 0,
					Place:            0,
					Tag:              0,
					Type:             0,
					ContactWayType:   "",
					ContactWay:       "",
					Status:           0,
				},
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ConfirmRequirementExist(tt.args.requirements)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfirmRequirementExist() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ConfirmRequirementExist() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCreatRequirement(t *testing.T) {
	Db.Init()
	defer Db.Close()
	type args struct {
		requirements Requirements
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				requirements: Requirements{
					RequirementId:    0,
					SenderSid:        "2019214300",
					Title:            "1",
					Content:          "2",
					PostTime:         "",
					Date:             0,
					TimeFrom:         3,
					TimeEnd:          4,
					RequirePeopleNum: 1,
					Place:            1,
					Tag:              4,
					Type:             4,
					ContactWayType:   "Q",
					ContactWay:       "QQQ",
					Status:           0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreatRequirement(tt.args.requirements); (err != nil) != tt.wantErr {
				t.Errorf("CreatRequirement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreatUser(t *testing.T) {
	type args struct {
		tmpUser UserInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				tmpUser: UserInfo{
					Sid:      "",
					NickName: "",
					College:  "",
					Gender:   "",
					Grade:    "",
					Portrait: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreatUser(tt.args.tmpUser); (err != nil) != tt.wantErr {
				t.Errorf("CreatUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFindUser(t *testing.T) {
	type args struct {
		uid string
	}
	tests := []struct {
		name    string
		args    args
		want    UserInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUser(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInfoFromRequirementId(t *testing.T) {
	type args struct {
		requirementID int
	}
	tests := []struct {
		name    string
		args    args
		want    Requirements
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInfoFromRequirementId(tt.args.requirementID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfoFromRequirementId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfoFromRequirementId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistoryRequirementFind(t *testing.T) {
	type args struct {
		uid    string
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		want    []HistoryRequirement
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HistoryRequirementFind(tt.args.uid, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("HistoryRequirementFind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HistoryRequirementFind() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecordAction(t *testing.T) {
	type args struct {
		uid string
		num int
		t   string
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
			if err := RecordAction(tt.args.uid, tt.args.num, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("RecordAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReminderBox(t *testing.T) {
	type args struct {
		uid    string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    []ReminderInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReminderBox(tt.args.uid, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReminderBox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReminderBox() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReminderChangeStatus(t *testing.T) {
	type args struct {
		applicationId int
		sid           string
		type1         int
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ReminderChangeStatus(tt.args.applicationId, tt.args.sid, tt.args.type1)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReminderChangeStatus() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReminderChangeStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRequirementApply(t *testing.T) {
	type args struct {
		uid             string
		requirementId   int
		contractWayType string
		contractWay     string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RequirementApply(tt.args.uid, tt.args.requirementId, tt.args.contractWayType, tt.args.contractWay)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequirementApply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RequirementApply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequirementDelete(t *testing.T) {
	type args struct {
		requirementId int
		sid           string
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RequirementDelete(tt.args.requirementId, tt.args.sid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequirementDelete() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RequirementDelete() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRequirementFind(t *testing.T) {
	type args struct {
		type1    int
		sid      string
		date     int
		timeFrom int
		timeEnd  int
		tag      []int
		place    []int
		limit    int
		offset   int
	}
	tests := []struct {
		name    string
		args    args
		want    []requirementInSquare
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RequirementFind(tt.args.type1, tt.args.sid, tt.args.date, tt.args.timeFrom, tt.args.timeEnd, tt.args.tag, tt.args.place, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequirementFind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequirementFind() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequirementInfo(t *testing.T) {
	type args struct {
		requirementId int
	}
	tests := []struct {
		name    string
		args    args
		want    Requirement
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := RequirementInfo(tt.args.requirementId)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequirementInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequirementInfo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RequirementInfo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSolveApplication(t *testing.T) {
	type args struct {
		applicationId int
		status        int
		sid           string
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SolveApplication(tt.args.applicationId, tt.args.status, tt.args.sid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveApplication() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SolveApplication() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestVerifyInfo(t *testing.T) {
	type args struct {
		uid        string
		verifyItem string
		verifyInfo string
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
			if err := VerifyInfo(tt.args.uid, tt.args.verifyItem, tt.args.verifyInfo); (err != nil) != tt.wantErr {
				t.Errorf("VerifyInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestViewAllApplication(t *testing.T) {
	type args struct {
		uid    string
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		want    []ViewApplicationInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ViewAllApplication(tt.args.uid, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ViewAllApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ViewAllApplication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRandomNum(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRandomNum(); got != tt.want {
				t.Errorf("getRandomNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redPoint(t *testing.T) {
	type args struct {
		status int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redPoint(tt.args.status); got != tt.want {
				t.Errorf("redPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
