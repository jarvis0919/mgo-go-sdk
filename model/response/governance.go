package response

type DelegatedStakeInfo struct {
	StakedMgoId       string `json:"stakedMgoId"`
	StakeRequestEpoch string `json:"stakeRequestEpoch"`
	StakeActiveEpoch  string `json:"stakeActiveEpoch"`
	Principal         string `json:"principal"`
	Status            string `json:"status"`
	EstimatedReward   string `json:"estimatedReward"`
}

type DelegatedStakesResponse struct {
	ValidatorAddress string               `json:"validatorAddress"`
	StakingPool      string               `json:"stakingPool"`
	Stakes           []DelegatedStakeInfo `json:"stakes"`
}
