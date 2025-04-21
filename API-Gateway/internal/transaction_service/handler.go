package transaction_service

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/gin-gonic/gin"

    mompb "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto"
    "github.com/JeroZp/gRPC-MOM/API-Gateway/internal/mom_service"
)

type transferReq struct {
    FromId string `json:"from_id" binding:"required"`
    ToId   string `json:"to_id"   binding:"required"`
    Amount int32  `json:"amount"  binding:"required,gt=0"`
}

// Transfer encola una solicitud de transferencia en MOM (“tx_ops”)
func Transfer(c *gin.Context) {
    var req transferReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    data, _ := json.Marshal(req)
    resp, err := mom_service.MOMClient.Publish(context.Background(), &mompb.PublishRequest{
        Queue: "tx_ops",
        Data:  data,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusAccepted, gin.H{"enqueued_id": resp.GetId()})
}