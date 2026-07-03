package unit

import (
	"testing"

	"blog/internal/model/dto/response"
)

func TestNewPageResponse(t *testing.T) {
	tests := []struct {
		name      string
		list      interface{}
		total     int64
		page      int
		size      int
		wantTotal int64
		wantPages int
	}{
		{
			name:      "正常分页",
			list:      []string{"a", "b"},
			total:     10,
			page:      1,
			size:      2,
			wantTotal: 10,
			wantPages: 5,
		},
		{
			name:      "空列表",
			list:      []string{},
			total:     0,
			page:      1,
			size:      10,
			wantTotal: 0,
			wantPages: 0,
		},
		{
			name:      "单页",
			list:      []int{1, 2, 3},
			total:     3,
			page:      1,
			size:      10,
			wantTotal: 3,
			wantPages: 1,
		},
		{
			name:      "最后一页不完整",
			list:      []int{11},
			total:     11,
			page:      2,
			size:      10,
			wantTotal: 11,
			wantPages: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := response.NewPageResponse(tt.list, tt.total, tt.page, tt.size)
			if resp.Total != tt.wantTotal {
				t.Errorf("Total = %v, want %v", resp.Total, tt.wantTotal)
			}
			if resp.TotalPage != tt.wantPages {
				t.Errorf("TotalPage = %v, want %v", resp.TotalPage, tt.wantPages)
			}
			if resp.Page != tt.page {
				t.Errorf("Page = %v, want %v", resp.Page, tt.page)
			}
			if resp.Size != tt.size {
				t.Errorf("Size = %v, want %v", resp.Size, tt.size)
			}
		})
	}
}

func TestNewPageResponse_TotalPageCalculation(t *testing.T) {
	// 整除情况
	resp := response.NewPageResponse(nil, 20, 1, 10)
	if resp.TotalPage != 2 {
		t.Errorf("TotalPage = %v, want 2", resp.TotalPage)
	}

	// 非整除情况
	resp = response.NewPageResponse(nil, 21, 1, 10)
	if resp.TotalPage != 3 {
		t.Errorf("TotalPage = %v, want 3", resp.TotalPage)
	}
}
