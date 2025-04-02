package metric

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockhand "github.com/SrVariable/mongo-exporter/api/metric/mock"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
	mockrepo "github.com/SrVariable/mongo-exporter/internal/metric/mock/repository"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRam(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ram := value_object.Ram{
		Resident: domain.Metric{Value: int32(10000)},
		Virtual:  domain.Metric{Value: int32(50000)},
	}

	repo := mockrepo.NewMockRepository(nil, nil, nil, &ram)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/v1/metrics/ram", nil)
	mockhand.GetRamHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := ram
	var got value_object.Ram
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want.Resident.Value, int32(got.Resident.Value.(float64)))
	assert.Equal(t, want.Virtual.Value, int32(got.Virtual.Value.(float64)))
}
