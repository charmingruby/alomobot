# Pools and Favs Cup

`/pool flip` deals entries without replacement into named groups:

```text
/pool flip groups:G,F pool:Favs size:6
/pool flip groups:Alice,Bob custom:larvitar,excadrill,scizor,gligar size:2
```

- `groups`: required comma-separated identifiers for the recipients.
- `pool`: a choice of registered pools.
- `custom`: comma-separated PokéAPI names, for this execution only.
- `size`: entries per group; defaults to 6.

Provide exactly one of `pool` or `custom`. Group identifiers are unique
(case-insensitive) and support letters, digits, hyphens and underscores, up to
32 characters. Insufficient entries fail rather than silently repeating Pokémon.
Repeated Pokémon in the source count as separate copies, including when
different contributors share a favorite. These copies can land in the same group.

Registered pools are listed in `internal/core/pool/pool.go`, with names defined
by `PoolName` constants such as `FavsPool`. Favs is defined in
`internal/core/pool/favs_pool.go`. Each pool has a name and
keyed subpools, so contributor identities remain available in the definition.
All subpools are combined for a draw; recipient group names do not select or
restrict the contributing subpools.

Subpools use `[]pkmn.Pokemon` with only `Name` and `ID` populated. Registered
entries are resolved by ID before drawing to load types and sprites; custom
entries are resolved by name, using the same `GetPokemon` lookup function.
Use the PokéAPI Pokémon ID for regional forms in registered pools
(e.g. `slowpoke-galar`: 10164, `ninetales-alola`: 10104), not the species'
National Dex number, which would resolve the default form.

`Favs` starts with:

- **G**: Incineroar, Scizor, Gligar, Slowpoke, Galarian Slowpoke, Alomomola, Reshiram.
- **F**: Tyranitar, Garchomp, Blaziken, Excadrill, Alolan Ninetales, Rayquaza.

Use the canonical form names `slowpoke-galar` and `ninetales-alola` in custom
pools. The two Slowpoke forms are separate entries. Favs has 13 entries:
two groups of six leave one unused entry.

The Favs Cup is registered for `/cup rules`, `/cup flip` and league draws.
Its rules refer to `/pool flip groups:G,F pool:Favs size:6`.
