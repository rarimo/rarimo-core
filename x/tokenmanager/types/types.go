package types

const (
	EventTypeCollectionCreated = "collection_created"
	EventTypeCollectionRemoved = "collection_removed"

	EventTypeCollectionDataCreated = "collection_data_created"
	EventTypeCollectionDataUpdated = "collection_data_updated"
	EventTypeCollectionDataRemoved = "collection_data_removed"

	EventTypeItemCreated = "item_created"
	EventTypeItemRemoved = "item_removed"

	EventTypeSeedCreated = "seed_created"
	EventTypeSeedRemoved = "seed_removed"

	EventTypeOnChainItemCreated = "on_chain_item_created"
	EventTypeOnChainItemRemoved = "on_chain_item_removed"

	AttributeKeyCollectionIndex     = "collection_index"
	AttributeKeyCollectionDataChain = "collection_data.chain"

	AttributeKeyItemIndex        = "item_index"
	AttributeKeyOnChainItemChain = "on_chain_item.chain"

	AttributeKeySeed = "seed"
)
