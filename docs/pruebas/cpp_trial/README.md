-- Terminal

cmake -G "MinGW Makefiles" ..
MinGW32-make

-- CMakeLists.txt

include(FetchContent)
FetchContent_Declare(cpr GIT_REPOSITORY https://github.com/libcpr/cpr.git
                         GIT_TAG 3020c34ae2b732121f37433e61599c34535e68a8) # The commit hash for 1.10.x. Replace with the latest from: https://github.com/libcpr/cpr/releases
FetchContent_MakeAvailable(cpr)


-- Idea Estructuras

Fixture {

// Estructura para "Ambos hacen gol"
GolesPorEquipo   := DictTrie<Equipo, DictTrie<idPartido, int>>
GolesRecibidos   := DictTrie<Equipo, DictTrie<idPartido, int>>
AmbosHicieronGol := DictrTrie<idPartido, int>

// Estructura para "Total X por partido"
TotalXPorPartido := ListaEnlazada<Struct<equipo1: int, equipo2: int>>
TotalXPorEquipo  := DictTrie<Equipo, DictTrie<idPartido, int>>

}
