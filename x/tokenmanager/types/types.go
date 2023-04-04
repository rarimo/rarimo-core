package types

const (
	EventTypeCollectionCreated = "collection_created" // emitted only from proposals
	EventTypeCollectionRemoved = "collection_removed" // emitted only from proposals

	EventTypeCollectionDataCreated = "collection_data_created" // emitted only from proposals
	EventTypeCollectionDataUpdated = "collection_data_updated" // emitted only from proposals
	EventTypeCollectionDataRemoved = "collection_data_removed" // emitted only from proposals

	EventTypeItemCreated = "item_created" // emitted from both transactions and proposals
	EventTypeItemRemoved = "item_removed" // emitted only from proposals

	EventTypeSeedCreated = "seed_created" // emitted from both transactions and proposals
	EventTypeSeedRemoved = "seed_removed" // emitted only from proposals

	EventTypeOnChainItemCreated = "on_chain_item_created" // emitted from both transactions and proposals
	EventTypeOnChainItemRemoved = "on_chain_item_removed" // emitted only from proposals

	AttributeKeyCollectionIndex     = "collection_index"
	AttributeKeyCollectionDataChain = "collection_data.chain"

	AttributeKeyItemIndex        = "item_index"
	AttributeKeyOnChainItemChain = "on_chain_item.chain"

	AttributeKeySeed = "seed"
)
