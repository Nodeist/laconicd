package simulation

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/cerc-io/laconicd/x/nameservice/types"
)

// RandomizedGenState generates a random GenesisState
func RandomizedGenState(simState *module.SimulationState) {
	nameserviceGenesis := types.DefaultGenesisState()

	bz, err := json.MarshalIndent(nameserviceGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(nameserviceGenesis)
}
