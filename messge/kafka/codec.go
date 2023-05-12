package confluentkafka

//
// import (
// 	"context"
// 	"fmt"
// 	"strconv"
// 	"time"
//
// 	confluentkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
// 	"github.com/go-leo/gox/slicex"
//
// 	"codeup.aliyun.com/qimao/leo/stream"
// )
//
// func DefaultMessageEncoder(topic string, msg stream.Message) (*confluentkafka.Message, error) {
// 	header := msg.Header()
//
// 	// partition
// 	partition := confluentkafka.PartitionAny
// 	if v := header.Get(PartitionKey); len(v) > 0 {
// 		intV, err := strconv.Atoi(v[0])
// 		if err != nil {
// 			return nil, err
// 		}
// 		partition = int32(intV)
// 	}
// 	header.Delete(PartitionKey)
//
// 	// key
// 	var key string
// 	if v := header.Get(KeyKey); len(v) > 0 {
// 		key = v[0]
// 	}
// 	header.Delete(KeyKey)
//
// 	opaque := header.Get(OpaqueKey)
// 	header.Delete(OpaqueKey)
//
// 	// other header
// 	kafkaHeaders := make([]confluentkafka.Header, 0, header.Len())
// 	header.Range(func(key string, values []string) {
// 		for _, value := range values {
// 			kafkaHeaders = append(kafkaHeaders, confluentkafka.Header{Key: key, Value: []byte(value)})
// 		}
// 	})
//
// 	kafkaMsg := &confluentkafka.Message{
// 		TopicPartition: confluentkafka.TopicPartition{
// 			Topic:     &topic,
// 			Partition: partition,
// 		},
// 		Value:         msg.Payload(),
// 		Key:           []byte(key),
// 		Timestamp:     time.Now(),
// 		TimestampType: confluentkafka.TimestampCreateTime,
// 		Opaque:        opaque,
// 		Headers:       kafkaHeaders,
// 	}
// 	return kafkaMsg, nil
// }
//
// func DefaultMessageDecoder(ctx context.Context, kafkaMsg *confluentkafka.Message) (stream.Message, error) {
// 	header := stream.Header{}
// 	if slicex.IsNotEmpty(kafkaMsg.Key) {
// 		header.Set(KeyKey, string(kafkaMsg.Key))
// 	}
// 	if !kafkaMsg.Timestamp.IsZero() {
// 		header.Set(TimestampKey, strconv.FormatInt(kafkaMsg.Timestamp.Unix(), 10))
// 		header.Set(TimestampTypeKey, kafkaMsg.TimestampType.String())
// 	}
// 	if kafkaMsg.Opaque != nil {
// 		header.Set(OpaqueKey, fmt.Sprint(kafkaMsg.Opaque))
// 	}
//
// 	for _, h := range kafkaMsg.Headers {
// 		header.Append(h.Key, string(h.Value))
// 	}
//
// 	topicPartition := kafkaMsg.TopicPartition
// 	header.Set(PartitionKey, strconv.FormatInt(int64(topicPartition.Partition), 10))
// 	header.Set(OffsetKey, strconv.FormatInt(int64(topicPartition.Offset), 10))
// 	if topicPartition.Topic != nil {
// 		header.Set(TopicKey, *topicPartition.Topic)
// 	}
// 	if topicPartition.Metadata != nil {
// 		header.Set(MetadataKey, *topicPartition.Metadata)
// 	}
// 	if topicPartition.Error != nil {
// 		header.Set(ErrorKey, topicPartition.Error.Error())
// 	}
// 	msg := stream.NewMessageBuilder().
// 		SetPayload(kafkaMsg.Value).
// 		SetHeader(header).
// 		Build(ctx)
// 	return msg, nil
// }
