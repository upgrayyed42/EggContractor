#!/usr/bin/env python3

import json
import pathlib

HERE = pathlib.Path(__file__).parent
DATAFILE = HERE / "../port/wasm/_common/eiafx/eiafx-data.json"
GOFILE = HERE / "effect_numbers.go"

NAME = {
    0: "api.ArtifactSpec_LUNAR_TOTEM",
    3: "api.ArtifactSpec_NEODYMIUM_MEDALLION",
    4: "api.ArtifactSpec_BEAK_OF_MIDAS",
    5: "api.ArtifactSpec_LIGHT_OF_EGGENDIL",
    6: "api.ArtifactSpec_DEMETERS_NECKLACE",
    7: "api.ArtifactSpec_VIAL_MARTIAN_DUST",
    8: "api.ArtifactSpec_ORNATE_GUSSET",
    9: "api.ArtifactSpec_THE_CHALICE",
    10: "api.ArtifactSpec_BOOK_OF_BASAN",
    11: "api.ArtifactSpec_PHOENIX_FEATHER",
    12: "api.ArtifactSpec_TUNGSTEN_ANKH",
    21: "api.ArtifactSpec_AURELIAN_BROOCH",
    22: "api.ArtifactSpec_CARVED_RAINSTICK",
    23: "api.ArtifactSpec_PUZZLE_CUBE",
    24: "api.ArtifactSpec_QUANTUM_METRONOME",
    25: "api.ArtifactSpec_SHIP_IN_A_BOTTLE",
    26: "api.ArtifactSpec_TACHYON_DEFLECTOR",
    27: "api.ArtifactSpec_INTERSTELLAR_COMPASS",
    28: "api.ArtifactSpec_DILITHIUM_MONOCLE",
    29: "api.ArtifactSpec_TITANIUM_ACTUATOR",
    30: "api.ArtifactSpec_MERCURYS_LENS",
    1: "api.ArtifactSpec_TACHYON_STONE",
    31: "api.ArtifactSpec_DILITHIUM_STONE",
    32: "api.ArtifactSpec_SHELL_STONE",
    33: "api.ArtifactSpec_LUNAR_STONE",
    34: "api.ArtifactSpec_SOUL_STONE",
    39: "api.ArtifactSpec_PROPHECY_STONE",
    36: "api.ArtifactSpec_QUANTUM_STONE",
    37: "api.ArtifactSpec_TERRA_STONE",
    38: "api.ArtifactSpec_LIFE_STONE",
    40: "api.ArtifactSpec_CLARITY_STONE",
    13: "api.ArtifactSpec_EXTRATERRESTRIAL_ALUMINUM",
    14: "api.ArtifactSpec_ANCIENT_TUNGSTEN",
    15: "api.ArtifactSpec_SPACE_ROCKS",
    16: "api.ArtifactSpec_ALIEN_WOOD",
    17: "api.ArtifactSpec_GOLD_METEORITE",
    18: "api.ArtifactSpec_TAU_CETI_GEODE",
    19: "api.ArtifactSpec_CENTAURIAN_STEEL",
    20: "api.ArtifactSpec_ERIDANI_FEATHER",
    35: "api.ArtifactSpec_DRONE_PARTS",
    41: "api.ArtifactSpec_CELESTIAL_BRONZE",
    42: "api.ArtifactSpec_LALANDE_HIDE",
    43: "api.ArtifactSpec_SOLAR_TITANIUM",
    2: "api.ArtifactSpec_TACHYON_STONE_FRAGMENT",
    44: "api.ArtifactSpec_DILITHIUM_STONE_FRAGMENT",
    45: "api.ArtifactSpec_SHELL_STONE_FRAGMENT",
    46: "api.ArtifactSpec_LUNAR_STONE_FRAGMENT",
    47: "api.ArtifactSpec_SOUL_STONE_FRAGMENT",
    48: "api.ArtifactSpec_PROPHECY_STONE_FRAGMENT",
    49: "api.ArtifactSpec_QUANTUM_STONE_FRAGMENT",
    50: "api.ArtifactSpec_TERRA_STONE_FRAGMENT",
    51: "api.ArtifactSpec_LIFE_STONE_FRAGMENT",
    52: "api.ArtifactSpec_CLARITY_STONE_FRAGMENT",
    10000: "api.ArtifactSpec_UNKNOWN",
}

LEVEL = {
    0: "api.ArtifactSpec_INFERIOR",
    1: "api.ArtifactSpec_LESSER",
    2: "api.ArtifactSpec_NORMAL",
    3: "api.ArtifactSpec_GREATER",
    4: "api.ArtifactSpec_SUPERIOR",
}

RARITY = {
    0: "api.ArtifactSpec_COMMON",
    1: "api.ArtifactSpec_RARE",
    2: "api.ArtifactSpec_EPIC",
    3: "api.ArtifactSpec_LEGENDARY",
}

"""
var _effectDeltas = map[item]float64{
	{
		Name:   0,
		Level:  0,
		Rarity: 0,
	}: 1.0,
}
"""


def main():
    with GOFILE.open("w") as fout:
        fout.write(
            """\
// Code generated by "python3 effect_numbers.py"; DO NOT EDIT.

package artifacts

import "github.com/fanaticscripter/EggContractor/api"

var _effectDeltas = map[item]float64{
"""
        )

        with DATAFILE.open() as fin:
            data = json.load(fin)
        for family in data["artifact_families"]:
            for tier in family["tiers"]:
                if tier["effects"] is None:
                    continue
                name = NAME[tier["afx_id"]]
                level = LEVEL[tier["afx_level"]]
                for rarity_effect in tier["effects"]:
                    rarity = RARITY[rarity_effect["afx_rarity"]]
                    delta = effect_delta(rarity_effect["effect_size"])
                    fout.write(
                        f"""\
	{{
		Name:   {name},
		Level:  {level},
		Rarity: {rarity},
	}}: {delta},
"""
                    )

        fout.write(
            """\
}
"""
        )


def effect_delta(effect_size: str) -> float:
    s = effect_size
    if s == "Guaranteed":
        return 0
    if s[-1] == "x":
        return float(s[:-1]) - 1
    if s[-1] == "%":
        return float(s[:-1]) * 0.01
    return float(s)


if __name__ == "__main__":
    main()