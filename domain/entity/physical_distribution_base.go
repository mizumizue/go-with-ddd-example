package entity

// 物流拠点
type PhysicalDistributionBase struct {
}

// 出庫
func (base *PhysicalDistributionBase) Ship(baggage *Baggage) *Baggage {
	panic("impl")
}

// 入庫
func (base *PhysicalDistributionBase) Receive(baggage *Baggage) {
	panic("impl")
}

// 出庫と入庫をセットで行うことを保証する為に、Transportを定義する
//func (base *PhysicalDistributionBase) Transport(to *PhysicalDistributionBase, baggage *Baggage) {
//	base.Ship(baggage)
//	to.Receive(baggage)
//
//	// 配送記録も必要？
//	panic("impl")
//}

// from から to へ輸送する為に拠点オブジェクトが全部責任を負うのは違和感がある？
// 別サービスとして定義する必要性
