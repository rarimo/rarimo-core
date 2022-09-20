# Go Merkle tree light implementation

This implementation will construct the Merkle tree where parent hash are hash of concatenation of its leafs,
but firstly will go lexicographically larger hash. Also, to check the path ypu should concatenate hashes in
lexicographically order.