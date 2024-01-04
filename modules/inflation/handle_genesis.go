package inflation

import (
	"encoding/json"

	tmtypes "github.com/tendermint/tendermint/types"

	inflationtypes "github.com/MonikaCat/em-ledger/x/inflation/types"
	"github.com/rs/zerolog/log"

	"github.com/forbole/bdjuno/v4/types"
)

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(genesisDoc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "inflation").Msg("parsing genesis")

	// Read the genesis state
	var genState inflationtypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[inflationtypes.ModuleName], &genState)
	if err != nil {
		return err
	}

	return m.db.SaveEMoneyInflation(types.NewEMoneyInflation(genState.InflationState, genesisDoc.InitialHeight))
}