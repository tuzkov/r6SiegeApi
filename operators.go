package r6

import (
	"encoding/json"
	"log"
	"strings"
)

const operatorJSON = `
{
  "smoke": {
    "id": "smoke",
    "category": "def",
    "name": {
      "oasisId": "62253"
    },
    "ctu": {
      "oasisId": "62015"
    },
    "index": "2:1",
    "figure": {
      "small": "assets/images/small-smoke.2726e30872.png",
      "large": "assets/images/large-smoke.1bf9006654.png"
    },
    "mask": "assets/images/mask-smoke.c84491688f.png",
    "badge": "assets/images/badge-smoke.874e98880d.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_smoke_poisongaskill:2:1",
        "label": {
          "oasisId": "194660"
        }
      },
      "pve": {
        "statisticId": "operatorpve_smoke_poisongaskill:2:1",
        "label": {
          "oasisId": "194660"
        }
      }
    }
  },
  "mute": {
    "id": "mute",
    "category": "def",
    "name": {
      "oasisId": "62252"
    },
    "ctu": {
      "oasisId": "62015"
    },
    "index": "3:1",
    "figure": {
      "small": "assets/images/small-mute.96a6566576.png",
      "large": "assets/images/large-mute.ae51429f4f.png"
    },
    "mask": "assets/images/mask-mute.d9917f856f.png",
    "badge": "assets/images/badge-mute.3e4f2b0170.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_mute_gadgetjammed:3:1",
        "label": {
          "oasisId": "194664"
        }
      },
      "pve": {
        "statisticId": "operatorpve_mute_gadgetjammed:3:1",
        "label": {
          "oasisId": "194664"
        }
      }
    }
  },
  "sledge": {
    "id": "sledge",
    "category": "atk",
    "name": {
      "oasisId": "62245"
    },
    "ctu": {
      "oasisId": "62015"
    },
    "index": "4:1",
    "figure": {
      "small": "assets/images/small-sledge.9ade04d351.png",
      "large": "assets/images/large-sledge.832f6c6b3c.png"
    },
    "mask": "assets/images/mask-sledge.08fa548b40.png",
    "badge": "assets/images/badge-sledge.00141f9258.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_sledge_hammerhole:4:1",
        "label": {
          "oasisId": "194672"
        }
      },
      "pve": {
        "statisticId": "operatorpve_sledge_hammerhole:4:1",
        "label": {
          "oasisId": "194672"
        }
      }
    }
  },
  "thatcher": {
    "id": "thatcher",
    "category": "atk",
    "name": {
      "oasisId": "65143"
    },
    "ctu": {
      "oasisId": "62015"
    },
    "index": "5:1",
    "figure": {
      "small": "assets/images/small-thatcher.df2b7de48d.png",
      "large": "assets/images/large-thatcher.73132fcdbe.png"
    },
    "mask": "assets/images/mask-thatcher.b7c658b343.png",
    "badge": "assets/images/badge-thatcher.b1cac8e7c0.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_thatcher_gadgetdestroywithemp:5:1",
        "label": {
          "oasisId": "194673"
        }
      },
      "pve": {
        "statisticId": "operatorpve_thatcher_gadgetdestroywithemp:5:1",
        "label": {
          "oasisId": "194673"
        }
      }
    }
  },
  "castle": {
    "id": "castle",
    "category": "def",
    "name": {
      "oasisId": "62248"
    },
    "ctu": {
      "oasisId": "62183"
    },
    "index": "2:2",
    "figure": {
      "small": "assets/images/small-castle.cd963832fb.png",
      "large": "assets/images/large-castle.6a1677f53f.png"
    },
    "mask": "assets/images/mask-castle.08abe353e2.png",
    "badge": "assets/images/badge-castle.378f8f4e24.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_castle_kevlarbarricadedeployed:2:2",
        "label": {
          "oasisId": "194663"
        }
      },
      "pve": {
        "statisticId": "operatorpve_castle_kevlarbarricadedeployed:2:2",
        "label": {
          "oasisId": "194663"
        }
      }
    }
  },
  "ash": {
    "id": "ash",
    "category": "atk",
    "name": {
      "oasisId": "62246"
    },
    "ctu": {
      "oasisId": "62183"
    },
    "index": "3:2",
    "figure": {
      "small": "assets/images/small-ash.692668cc8a.png",
      "large": "assets/images/large-ash.9d28aebe23.png"
    },
    "mask": "assets/images/mask-ash.ad5c86d093.png",
    "badge": "assets/images/badge-ash.16913d82e3.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_ash_bonfirewallbreached:3:2",
        "label": {
          "oasisId": "194677"
        }
      },
      "pve": {
        "statisticId": "operatorpve_ash_bonfirewallbreached:3:2",
        "label": {
          "oasisId": "194677"
        }
      }
    }
  },
  "pulse": {
    "id": "pulse",
    "category": "def",
    "name": {
      "oasisId": "62250"
    },
    "ctu": {
      "oasisId": "62183"
    },
    "index": "4:2",
    "figure": {
      "small": "assets/images/small-pulse.927fa707a4.png",
      "large": "assets/images/large-pulse.30ab3682ce.png"
    },
    "mask": "assets/images/mask-pulse.0b2985b7ea.png",
    "badge": "assets/images/badge-pulse.9de627c54e.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_pulse_heartbeatspot:4:2",
        "label": {
          "oasisId": "194665"
        }
      },
      "pve": {
        "statisticId": "operatorpve_pulse_heartbeatspot:4:2",
        "label": {
          "oasisId": "194665"
        }
      }
    }
  },
  "thermite": {
    "id": "thermite",
    "category": "atk",
    "name": {
      "oasisId": "62247"
    },
    "ctu": {
      "oasisId": "62183"
    },
    "index": "5:2",
    "figure": {
      "small": "assets/images/small-thermite.f9441d1f1a.png",
      "large": "assets/images/large-thermite.e973bb0498.png"
    },
    "mask": "assets/images/mask-thermite.a46a67f0b5.png",
    "badge": "assets/images/badge-thermite.9010fa3311.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_thermite_reinforcementbreached:5:2",
        "label": {
          "oasisId": "194681"
        }
      },
      "pve": {
        "statisticId": "operatorpve_thermite_reinforcementbreached:5:2",
        "label": {
          "oasisId": "194681"
        }
      }
    }
  },
  "doc": {
    "id": "doc",
    "category": "def",
    "name": {
      "oasisId": "62251"
    },
    "ctu": {
      "oasisId": "62016"
    },
    "index": "2:3",
    "figure": {
      "small": "assets/images/small-doc.2cc6664a0e.png",
      "large": "assets/images/large-doc.0b0321eb2f.png"
    },
    "mask": "assets/images/mask-doc.951579fcec.png",
    "badge": "assets/images/badge-doc.29fe751bea.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_doc_teammaterevive:2:3",
        "label": {
          "oasisId": "194666"
        }
      },
      "pve": {
        "statisticId": "operatorpve_doc_teammaterevive:2:3",
        "label": {
          "oasisId": "194666"
        }
      }
    }
  },
  "rook": {
    "id": "rook",
    "category": "def",
    "name": {
      "oasisId": "62249"
    },
    "ctu": {
      "oasisId": "62016"
    },
    "index": "3:3",
    "figure": {
      "small": "assets/images/small-rook.55b6bc4c7e.png",
      "large": "assets/images/large-rook.eed2777a17.png"
    },
    "mask": "assets/images/mask-rook.be9b0ac99a.png",
    "badge": "assets/images/badge-rook.eb954a4e67.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_rook_armortakenteammate:3:3",
        "label": {
          "oasisId": "194667"
        }
      },
      "pve": {
        "statisticId": "operatorpve_rook_armortakenteammate:3:3",
        "label": {
          "oasisId": "194667"
        }
      }
    }
  },
  "twitch": {
    "id": "twitch",
    "category": "atk",
    "name": {
      "oasisId": "65156"
    },
    "ctu": {
      "oasisId": "62016"
    },
    "index": "4:3",
    "figure": {
      "small": "assets/images/small-twitch.0421265dc4.png",
      "large": "assets/images/large-twitch.84ad765cee.png"
    },
    "mask": "assets/images/mask-twitch.715147fcdf.png",
    "badge": "assets/images/badge-twitch.83cbfa9789.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_twitch_gadgetdestroybyshockdrone:4:3",
        "label": {
          "oasisId": "194686"
        }
      },
      "pve": {
        "statisticId": "operatorpve_twitch_gadgetdestroybyshockdrone:4:3",
        "label": {
          "oasisId": "194686"
        }
      }
    }
  },
  "montagne": {
    "id": "montagne",
    "category": "atk",
    "name": {
      "oasisId": "65159"
    },
    "ctu": {
      "oasisId": "62016"
    },
    "index": "5:3",
    "figure": {
      "small": "assets/images/small-montagne.0bc5133625.png",
      "large": "assets/images/large-montagne.d48591cc9f.png"
    },
    "mask": "assets/images/mask-montagne.b392c01592.png",
    "badge": "assets/images/badge-montagne.2078ee847a.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_montagne_shieldblockdamage:5:3",
        "label": {
          "oasisId": "194688"
        }
      },
      "pve": {
        "statisticId": "operatorpve_montagne_shieldblockdamage:5:3",
        "label": {
          "oasisId": "194688"
        }
      }
    }
  },
  "glaz": {
    "id": "glaz",
    "category": "atk",
    "name": {
      "oasisId": "62242"
    },
    "ctu": {
      "oasisId": "621850"
    },
    "index": "2:4",
    "figure": {
      "small": "assets/images/small-glaz.a89bfa89dc.png",
      "large": "assets/images/large-glaz.8cd96a16f8.png"
    },
    "mask": "assets/images/mask-glaz.f3a01dea74.png",
    "badge": "assets/images/badge-glaz.43dd3bdfbc.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_glaz_sniperkill:2:4",
        "label": {
          "oasisId": "194689"
        }
      },
      "pve": {
        "statisticId": "operatorpve_glaz_sniperkill:2:4",
        "label": {
          "oasisId": "194689"
        }
      }
    }
  },
  "fuze": {
    "id": "fuze",
    "category": "atk",
    "name": {
      "oasisId": "65168"
    },
    "ctu": {
      "oasisId": "621850"
    },
    "index": "3:4",
    "figure": {
      "small": "assets/images/small-fuze.279743fadb.png",
      "large": "assets/images/large-fuze.5f8fb3ba84.png"
    },
    "mask": "assets/images/mask-fuze.090fe62680.png",
    "badge": "assets/images/badge-fuze.9e7e92226e.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_fuze_clusterchargekill:3:4",
        "label": {
          "oasisId": "194691"
        }
      },
      "pve": {
        "statisticId": "operatorpve_fuze_clusterchargekill:3:4",
        "label": {
          "oasisId": "194691"
        }
      }
    }
  },
  "kapkan": {
    "id": "kapkan",
    "category": "def",
    "name": {
      "oasisId": "65171"
    },
    "ctu": {
      "oasisId": "621850"
    },
    "index": "4:4",
    "figure": {
      "small": "assets/images/small-kapkan.b796e6065e.png",
      "large": "assets/images/large-kapkan.606a9fc0d3.png"
    },
    "mask": "assets/images/mask-kapkan.93f6064146.png",
    "badge": "assets/images/badge-kapkan.562d0701e7.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_kapkan_boobytrapkill:4:4",
        "label": {
          "oasisId": "194668"
        }
      },
      "pve": {
        "statisticId": "operatorpve_kapkan_boobytrapkill:4:4",
        "label": {
          "oasisId": "194668"
        }
      }
    }
  },
  "tachanka": {
    "id": "tachanka",
    "category": "def",
    "name": {
      "oasisId": "65174"
    },
    "ctu": {
      "oasisId": "621850"
    },
    "index": "5:4",
    "figure": {
      "small": "assets/images/small-tachanka.4c41f39c84.png",
      "large": "assets/images/large-tachanka.41caebce49.png"
    },
    "mask": "assets/images/mask-tachanka.a613ee385a.png",
    "badge": "assets/images/badge-tachanka.ae7943f00d.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_tachanka_turretkill:5:4",
        "label": {
          "oasisId": "194669"
        }
      },
      "pve": {
        "statisticId": "operatorpve_tachanka_turretkill:5:4",
        "label": {
          "oasisId": "194669"
        }
      }
    }
  },
  "blitz": {
    "id": "blitz",
    "category": "atk",
    "name": {
      "oasisId": "62243"
    },
    "ctu": {
      "oasisId": "62184"
    },
    "index": "2:5",
    "figure": {
      "small": "assets/images/small-blitz.0553bf11e0.png",
      "large": "assets/images/large-blitz.2e74d4b8ac.png"
    },
    "mask": "assets/images/mask-blitz.bcde0a3313.png",
    "badge": "assets/images/badge-blitz.cd45df08f6.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_blitz_flashedenemy:2:5",
        "label": {
          "oasisId": "195161"
        }
      },
      "pve": {
        "statisticId": "operatorpve_blitz_flashedenemy:2:5",
        "label": {
          "oasisId": "195161"
        }
      }
    }
  },
  "iq": {
    "id": "iq",
    "category": "atk",
    "name": {
      "oasisId": "113052"
    },
    "ctu": {
      "oasisId": "62184"
    },
    "index": "3:5",
    "figure": {
      "small": "assets/images/small-iq.0d7885fbe8.png",
      "large": "assets/images/large-iq.22b58f3bd4.png"
    },
    "mask": "assets/images/mask-iq.855499c7b3.png",
    "badge": "assets/images/badge-iq.b1acee1a4c.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_iq_gadgetspotbyef:3:5",
        "label": {
          "oasisId": "194696"
        }
      },
      "pve": {
        "statisticId": "operatorpve_iq_gadgetspotbyef:3:5",
        "label": {
          "oasisId": "194696"
        }
      }
    }
  },
  "jager": {
    "id": "jager",
    "category": "def",
    "name": {
      "oasisId": "113049"
    },
    "ctu": {
      "oasisId": "62184"
    },
    "index": "4:5",
    "figure": {
      "small": "assets/images/small-jager.bc9087b7b2.png",
      "large": "assets/images/large-jager.edd9f74950.png"
    },
    "mask": "assets/images/mask-jager.58cba6d598.png",
    "badge": "assets/images/badge-jager.600b2773be.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_jager_gadgetdestroybycatcher:4:5",
        "label": {
          "oasisId": "194670"
        }
      },
      "pve": {
        "statisticId": "operatorpve_jager_gadgetdestroybycatcher:4:5",
        "label": {
          "oasisId": "194670"
        }
      }
    }
  },
  "bandit": {
    "id": "bandit",
    "category": "def",
    "name": {
      "oasisId": "65165"
    },
    "ctu": {
      "oasisId": "62184"
    },
    "index": "5:5",
    "figure": {
      "small": "assets/images/small-bandit.113ec809d3.png",
      "large": "assets/images/large-bandit.463ee8297b.png"
    },
    "mask": "assets/images/mask-bandit.8879a650ec.png",
    "badge": "assets/images/badge-bandit.385144d970.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_bandit_batterykill:5:5",
        "label": {
          "oasisId": "194671"
        }
      },
      "pve": {
        "statisticId": "operatorpve_bandit_batterykill:5:5",
        "label": {
          "oasisId": "194671"
        }
      }
    }
  },
  "buck": {
    "id": "buck",
    "category": "atk",
    "name": {
      "oasisId": "62244"
    },
    "ctu": {
      "oasisId": "192111"
    },
    "index": "2:6",
    "figure": {
      "small": "assets/images/small-buck.6b66cd57a3.png",
      "large": "assets/images/large-buck.78712d24f8.png"
    },
    "mask": "assets/images/mask-buck.e9bc3630cf.png",
    "badge": "assets/images/badge-buck.2fc3e09779.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_buck_kill:2:6",
        "label": {
          "oasisId": "200102"
        }
      },
      "pve": {
        "statisticId": "operatorpve_buck_kill:2:6",
        "label": {
          "oasisId": "200102"
        }
      }
    }
  },
  "frost": {
    "id": "frost",
    "category": "def",
    "name": {
      "oasisId": "186159"
    },
    "ctu": {
      "oasisId": "192111"
    },
    "index": "3:6",
    "figure": {
      "small": "assets/images/small-frost.a7f18f5b86.png",
      "large": "assets/images/large-frost.f4325d10e2.png"
    },
    "mask": "assets/images/mask-frost.51971e4dda.png",
    "badge": "assets/images/badge-frost.e5362220a9.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_frost_dbno:3:6",
        "label": {
          "oasisId": "200104"
        }
      },
      "pve": {
        "statisticId": "operatorpve_frost_beartrap_kill:3:6",
        "label": {
          "oasisId": "200103"
        }
      }
    }
  },
  "blackbeard": {
    "id": "blackbeard",
    "category": "atk",
    "name": {
      "oasisId": "187249"
    },
    "ctu": {
      "oasisId": "62186"
    },
    "index": "2:7",
    "figure": {
      "small": "assets/images/small-blackbeard.cba9e22d86.png",
      "large": "assets/images/large-blackbeard.2292a7911f.png"
    },
    "mask": "assets/images/mask-blackbeard.4dbca57284.png",
    "badge": "assets/images/badge-blackbeard.fccd7e2c03.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_blackbeard_gunshieldblockdamage:2:7",
        "label": {
          "oasisId": "202310"
        }
      },
      "pve": {
        "statisticId": "operatorpve_blackbeard_gunshieldblockdamage:2:7",
        "label": {
          "oasisId": "202310"
        }
      }
    }
  },
  "valkyrie": {
    "id": "valkyrie",
    "category": "def",
    "name": {
      "oasisId": "187252"
    },
    "ctu": {
      "oasisId": "62186"
    },
    "index": "3:7",
    "figure": {
      "small": "assets/images/small-valkyrie.b2df5baedd.png",
      "large": "assets/images/large-valkyrie.c1f143fb3c.png"
    },
    "mask": "assets/images/mask-valkyrie.c93357698e.png",
    "badge": "assets/images/badge-valkyrie.f87cb6bdc2.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_valkyrie_camdeployed:3:7",
        "label": {
          "oasisId": "202311"
        }
      },
      "pve": {
        "statisticId": "operatorpve_valkyrie_camdeployed:3:7",
        "label": {
          "oasisId": "202311"
        }
      }
    }
  },
  "capitao": {
    "id": "capitao",
    "category": "atk",
    "name": {
      "oasisId": "207674"
    },
    "ctu": {
      "oasisId": "207757"
    },
    "index": "2:8",
    "figure": {
      "small": "assets/images/small-capitao.31c21fd075.png",
      "large": "assets/images/large-capitao.984e75b759.png"
    },
    "mask": "assets/images/mask-capitao.f56d66af19.png",
    "badge": "assets/images/badge-capitao.6603e417c1.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_capitao_lethaldartkills:2:8",
        "label": {
          "oasisId": "207946"
        }
      },
      "pve": {
        "statisticId": "operatorpve_capitao_lethaldartkills:2:8",
        "label": {
          "oasisId": "207946"
        }
      }
    }
  },
  "caveira": {
    "id": "caveira",
    "category": "def",
    "name": {
      "oasisId": "207671"
    },
    "ctu": {
      "oasisId": "207757"
    },
    "index": "3:8",
    "figure": {
      "small": "assets/images/small-caveira.f7bb7af3be.png",
      "large": "assets/images/large-caveira.e4d82365c0.png"
    },
    "mask": "assets/images/mask-caveira.2971b2c5e0.png",
    "badge": "assets/images/badge-caveira.757e9259e4.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_caveira_interrogations:3:8",
        "label": {
          "oasisId": "207945"
        }
      },
      "pve": {
        "statisticId": "operatorpve_caveira_aikilledinstealth:3:8",
        "label": {
          "oasisId": "207952"
        }
      }
    }
  },
  "hibana": {
    "id": "hibana",
    "category": "atk",
    "name": {
      "oasisId": "209183"
    },
    "ctu": {
      "oasisId": "210241"
    },
    "index": "2:9",
    "figure": {
      "small": "assets/images/small-hibana.494345ca09.png",
      "large": "assets/images/large-hibana.d3ceb7759e.png"
    },
    "mask": "assets/images/mask-hibana.b41fbe46f3.png",
    "badge": "assets/images/badge-hibana.c2a8477d1b.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_hibana_detonate_projectile:2:9",
        "label": {
          "oasisId": "217323"
        }
      },
      "pve": {
        "statisticId": "operatorpve_hibana_detonate_projectile:2:9",
        "label": {
          "oasisId": "217323"
        }
      }
    }
  },
  "echo": {
    "id": "echo",
    "category": "def",
    "name": {
      "oasisId": "209180"
    },
    "ctu": {
      "oasisId": "210241"
    },
    "index": "3:9",
    "figure": {
      "small": "assets/images/small-echo.6c60bd2a15.png",
      "large": "assets/images/large-echo.592cfb343c.png"
    },
    "mask": "assets/images/mask-echo.43fff69139.png",
    "badge": "assets/images/badge-echo.a77c7d7eb5.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_echo_enemy_sonicburst_affected:3:9",
        "label": {
          "oasisId": "217324"
        }
      },
      "pve": {
        "statisticId": "operatorpve_echo_enemy_sonicburst_affected:3:9",
        "label": {
          "oasisId": "217324"
        }
      }
    }
  },
  "jackal": {
    "id": "jackal",
    "category": "atk",
    "name": {
      "oasisId": "222182"
    },
    "ctu": {
      "oasisId": "217493"
    },
    "index": "2:A",
    "figure": {
      "small": "assets/images/small-jackal.8c3e419115.png",
      "large": "assets/images/large-jackal.e7ec96e645.png"
    },
    "mask": "assets/images/mask-jackal.b183ae5448.png",
    "badge": "assets/images/badge-jackal.0326ca29ca.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_cazador_assist_kill:2:A",
        "label": {
          "oasisId": "229326"
        }
      },
      "pve": {
        "statisticId": "operatorpve_cazador_assist_kill:2:A",
        "label": {
          "oasisId": "229327"
        }
      }
    }
  },
  "mira": {
    "id": "mira",
    "category": "def",
    "name": {
      "oasisId": "222172"
    },
    "ctu": {
      "oasisId": "217493"
    },
    "index": "3:A",
    "figure": {
      "small": "assets/images/small-mira.4196bc073d.png",
      "large": "assets/images/large-mira.0c9e3bd8d8.png"
    },
    "mask": "assets/images/mask-mira.3d59212129.png",
    "badge": "assets/images/badge-mira.22fb72a5cc.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_black_mirror_gadget_deployed:3:A",
        "label": {
          "oasisId": "229328"
        }
      },
      "pve": {
        "statisticId": "operatorpve_black_mirror_gadget_deployed:3:A",
        "label": {
          "oasisId": "229329"
        }
      }
    }
  },
  "ying": {
    "id": "ying",
    "category": "atk",
    "name": {
      "oasisId": "239334"
    },
    "ctu": {
      "oasisId": "231289"
    },
    "index": "2:B",
    "figure": {
      "small": "assets/images/small-ying.6bec33ffd1.png",
      "large": "assets/images/large-ying.ce15b076b7.png"
    },
    "mask": "assets/images/mask-ying.0e91da74eb.png",
    "badge": "assets/images/badge-ying.b88be612bd.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_dazzler_gadget_detonate:2:B",
        "label": {
          "oasisId": "243475"
        }
      },
      "pve": {
        "statisticId": "operatorpve_dazzler_gadget_detonate:2:B",
        "label": {
          "oasisId": "243475"
        }
      }
    }
  },
  "lesion": {
    "id": "lesion",
    "category": "def",
    "name": {
      "oasisId": "239337"
    },
    "ctu": {
      "oasisId": "231289"
    },
    "index": "3:B",
    "figure": {
      "small": "assets/images/small-lesion.8165c6b9a4.png",
      "large": "assets/images/large-lesion.d48d6d3321.png"
    },
    "mask": "assets/images/mask-lesion.576ad9e46b.png",
    "badge": "assets/images/badge-lesion.07c3d352ca.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_caltrop_enemy_affected:3:B",
        "label": {
          "oasisId": "243474"
        }
      },
      "pve": {
        "statisticId": "operatorpve_caltrop_enemy_affected:3:B",
        "label": {
          "oasisId": "243474"
        }
      }
    }
  },
  "ela": {
    "id": "ela",
    "category": "def",
    "name": {
      "oasisId": "254595"
    },
    "ctu": {
      "oasisId": "254591"
    },
    "index": "2:C",
    "figure": {
      "small": "assets/images/small-ela.b8458c8473.png",
      "large": "assets/images/large-ela.a50c192d00.png"
    },
    "mask": "assets/images/mask-ela.eba6e6df45.png",
    "badge": "assets/images/badge-ela.63ec2d26e4.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_concussionmine_detonate:2:C",
        "label": {
          "oasisId": "254989"
        }
      },
      "pve": {
        "statisticId": "operatorpve_concussionmine_detonate:2:C",
        "label": {
          "oasisId": "254989"
        }
      }
    }
  },
  "zofia": {
    "id": "zofia",
    "category": "atk",
    "name": {
      "oasisId": "254606"
    },
    "ctu": {
      "oasisId": "254591"
    },
    "index": "3:C",
    "figure": {
      "small": "assets/images/small-zofia.28fa7ba766.png",
      "large": "assets/images/large-zofia.f9f7568bc4.png"
    },
    "mask": "assets/images/mask-zofia.29e5102fef.png",
    "badge": "assets/images/badge-zofia.2a892bf5dc.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_concussiongrenade_detonate:3:C",
        "label": {
          "oasisId": "254990"
        }
      },
      "pve": {
        "statisticId": "operatorpve_concussiongrenade_detonate:3:C",
        "label": {
          "oasisId": "254990"
        }
      }
    }
  },
  "dokkaebi": {
    "id": "dokkaebi",
    "category": "atk",
    "name": {
      "oasisId": "273652"
    },
    "ctu": {
      "oasisId": "273753"
    },
    "index": "2:D",
    "figure": {
      "small": "assets/images/small-dokkaebi.495c615562.png",
      "large": "assets/images/large-dokkaebi.17504ff442.png"
    },
    "mask": "assets/images/mask-dokkaebi.869d844585.png",
    "badge": "assets/images/badge-dokkaebi.2f83a34f88.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_phoneshacked:2:D",
        "label": {
          "oasisId": "283005"
        }
      },
      "pve": {
        "statisticId": "operatorpve_phoneshacked:2:D",
        "label": {
          "oasisId": "283005"
        }
      }
    }
  },
  "vigil": {
    "id": "vigil",
    "category": "def",
    "name": {
      "oasisId": "273663"
    },
    "ctu": {
      "oasisId": "273753"
    },
    "index": "3:D",
    "figure": {
      "small": "assets/images/small-vigil.0a3f4bbb3b.png",
      "large": "assets/images/large-vigil.00339c3468.png"
    },
    "mask": "assets/images/mask-vigil.8ef2f7017a.png",
    "badge": "assets/images/badge-vigil.4db5385b08.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_attackerdrone_diminishedrealitymode:3:D",
        "label": {
          "oasisId": "283004"
        }
      },
      "pve": {
        "statisticId": "operatorpve_attackerdrone_diminishedrealitymode:3:D",
        "label": {
          "oasisId": "283004"
        }
      }
    }
  },
  "lion": {
    "id": "lion",
    "category": "atk",
    "name": {
      "oasisId": "281463"
    },
    "ctu": {
      "oasisId": "282958"
    },
    "index": "3:E",
    "figure": {
      "small": "assets/images/small-lion.d9a0faed93.png",
      "large": "assets/images/large-lion.68296c0e60.png"
    },
    "mask": "assets/images/mask-lion.488ea9cca4.png",
    "badge": "assets/images/badge-lion.696370754d.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_tagger_tagdevice_spot:3:E",
        "label": {
          "oasisId": "282608"
        }
      },
      "pve": {
        "statisticId": "operatorpve_tagger_tagdevice_spot:3:E",
        "label": {
          "oasisId": "282608"
        }
      }
    }
  },
  "finka": {
    "id": "finka",
    "category": "atk",
    "name": {
      "oasisId": "281457"
    },
    "ctu": {
      "oasisId": "282958"
    },
    "index": "4:E",
    "figure": {
      "small": "assets/images/small-finka.53104c1345.png",
      "large": "assets/images/large-finka.46fb7c595e.png"
    },
    "mask": "assets/images/mask-finka.82e49f1a08.png",
    "badge": "assets/images/badge-finka.71d3a2432f.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_rush_adrenalinerush:4:E",
        "label": {
          "oasisId": "282609"
        }
      },
      "pve": {
        "statisticId": "operatorpve_rush_adrenalinerush:4:E",
        "label": {
          "oasisId": "282609"
        }
      }
    }
  },
  "maestro": {
    "id": "maestro",
    "category": "def",
    "name": {
      "oasisId": "285888"
    },
    "ctu": {
      "oasisId": "285859"
    },
    "index": "2:F",
    "figure": {
      "small": "assets/images/small-maestro.7f3f2bcb8b.png",
      "large": "assets/images/large-maestro.fc1e1753fc.png"
    },
    "mask": "assets/images/mask-maestro.210c8915d0.png",
    "badge": "assets/images/badge-maestro.b6cf7905ed.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_barrage_killswithturret:2:F",
        "label": {
          "oasisId": "288335"
        }
      },
      "pve": {
        "statisticId": "operatorpve_barrage_killswithturret:2:F",
        "label": {
          "oasisId": "288335"
        }
      }
    }
  },
  "alibi": {
    "id": "alibi",
    "category": "def",
    "name": {
      "oasisId": "282399"
    },
    "ctu": {
      "oasisId": "285859"
    },
    "index": "3:F",
    "figure": {
      "small": "assets/images/small-alibi.9b68753a5e.png",
      "large": "assets/images/large-alibi.3cf80a66b1.png"
    },
    "mask": "assets/images/mask-alibi.2f8c8bb326.png",
    "badge": "assets/images/badge-alibi.7fba8d3300.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_deceiver_revealedattackers:3:F",
        "label": {
          "oasisId": "282404"
        }
      },
      "pve": {
        "statisticId": "operatorpve_deceiver_revealedattackers:3:F",
        "label": {
          "oasisId": "282403"
        }
      }
    }
  },
  "maverick": {
    "id": "maverick",
    "category": "atk",
    "name": {
      "oasisId": "288504"
    },
    "ctu": {
      "oasisId": "289336"
    },
    "index": "2:10",
    "figure": {
      "small": "assets/images/small-maverick.9e685421f1.png",
      "large": "assets/images/large-maverick.2676adff62.png"
    },
    "mask": "assets/images/mask-maverick.d7db4066f7.png",
    "badge": "assets/images/badge-maverick.7eab7c75e7.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_maverick_wallbreached:2:10",
        "label": {
          "oasisId": "293475"
        }
      },
      "pve": {
        "statisticId": "operatorpve_maverick_wallbreached:2:10",
        "label": {
          "oasisId": "293475"
        }
      }
    }
  },
  "clash": {
    "id": "clash",
    "category": "def",
    "name": {
      "oasisId": "288526"
    },
    "ctu": {
      "oasisId": "289336"
    },
    "index": "3:10",
    "figure": {
      "small": "assets/images/small-clash.5d0756a523.png",
      "large": "assets/images/large-clash.fb3f6fb344.png"
    },
    "mask": "assets/images/mask-clash.118cccde64.png",
    "badge": "assets/images/badge-clash.133f243da3.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_clash_sloweddown:3:10",
        "label": {
          "oasisId": "293483"
        }
      },
      "pve": {
        "statisticId": "operatorpve_clash_sloweddown:3:10",
        "label": {
          "oasisId": "293483"
        }
      }
    }
  },
  "nomad": {
    "id": "nomad",
    "category": "atk",
    "name": {
      "oasisId": "293435"
    },
    "ctu": {
      "oasisId": "292860"
    },
    "index": "2:11",
    "figure": {
      "small": "assets/images/small-nomad.50642fb12f.png",
      "large": "assets/images/large-nomad.ec87176490.png"
    },
    "mask": "assets/images/mask-nomad.fdb16c874b.png",
    "badge": "assets/images/badge-nomad.dbd9a315fa.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_Nomad_Assist:2:11",
        "label": {
          "oasisId": "293935"
        }
      },
      "pve": {
        "statisticId": "operatorpve_Nomad_Assist:2:11",
        "label": {
          "oasisId": "293935"
        }
      }
    }
  },
  "kaid": {
    "id": "kaid",
    "category": "def",
    "name": {
      "oasisId": "293436"
    },
    "ctu": {
      "oasisId": "292860"
    },
    "index": "3:11",
    "figure": {
      "small": "assets/images/small-kaid.e38b6341ae.png",
      "large": "assets/images/large-kaid.d801cd6fa4.png"
    },
    "mask": "assets/images/mask-kaid.e8bb20882b.png",
    "badge": "assets/images/badge-kaid.ae2bfa7a5b.png",
    "uniqueStatistic": {
      "pvp": {
        "statisticId": "operatorpvp_Kaid_Electroclaw_Hatches:3:11",
        "label": {
          "oasisId": "294144"
        }
      },
      "pve": {
        "statisticId": "operatorpve_Kaid_Electroclaw_Hatches:3:11",
        "label": {
          "oasisId": "294144"
        }
      }
    }
  }
}
`

var (
	operatorSkillsTranslations = map[string]string{
		"death":                               "Смертей",
		"kills":                               "Убийств",
		"roundlost":                           "Проиграно раундов",
		"roundwon":                            "Выиграно раундов",
		"timeplayed":                          "Сыграно времени",
		"smoke_poisongaskill":                 "Убийств газом",
		"mute_gadgetjammed":                   "Отключенные джаммером устройства",
		"sledge_hammerhole":                   "Пробито кувалдой",
		"thatcher_gadgetdestroywithemp":       "Уничтожено устройств EMP",
		"castle_kevlarbarricadedeployed":      "Установлено бронепанелей",
		"ash_bonfirewallbreached":             "Использований гранатомета",
		"pulse_heartbeatspot":                 "Обнаружено сердцебиений",
		"thermite_reinforcementbreached":      "Пропилено стен",
		"doc_teammaterevive":                  "Поднято союзников",
		"rook_armortakenteammate":             "Поднято бронепластин",
		"twitch_gadgetdestroybyshockdrone":    "Уничтожено гаджетов шок-дроном",
		"montagne_shieldblockdamage":          "Заблокировано выстрелов",
		"glaz_sniperkill":                     "Убито в прицел",
		"fuze_clusterchargekill":              "Убийств кластером",
		"kapkan_boobytrapkill":                "Убийств капканом",
		"tachanka_turretkill":                 "Убийств с турели",
		"blitz_flashedenemy":                  "Ослеплено врагов",
		"iq_gadgetspotbyef":                   "Обнаружено гаджетов",
		"jager_gadgetdestroybycatcher":        "Сбито гранат",
		"bandit_batterykill":                  "Убийств электричеством",
		"buck_kill":                           "Убийств дробовиком",
		"frost_dbno":                          "Срабатываний силка",
		"blackbeard_gunshieldblockdamage":     "Поймано пуль щитом",
		"valkyrie_camdeployed":                "Установлено камер",
		"capitao_lethaldartkills":             "Сожжено врагов",
		"caveira_interrogations":              "Допросов",
		"hibana_detonate_projectile":          "Взорвано пеллетов",
		"echo_enemy_sonicburst_affected":      "Врагов оглушено",
		"cazador_assist_kill":                 "Помощь отслеживанием",
		"black_mirror_gadget_deployed":        "Установлено зеркал",
		"dazzler_gadget_detonate":             "Брошено канделл",
		"caltrop_enemy_affected":              "Отравлено врагов",
		"concussionmine_detonate":             "Оглушено миной",
		"concussiongrenade_detonate":          "Оглушено гранатой",
		"phoneshacked":                        "Взломано телефонов",
		"attackerdrone_diminishedrealitymode": "Обмануто дронов",
		"tagger_tagdevice_spot":               "Подсвечено врагов",
		"rush_adrenalinerush":                 "Пусков адреналина",
		"barrage_killswithturret":             "Убито турелью",
		"deceiver_revealedattackers":          "Обмануто врагов призмой",
		"maverick_wallbreached":               "Испорчено стен",
		"clash_sloweddown":                    "Замедленно врагов",
		"Nomad_Assist":                        "Активированные заряды Airjab",
		"Kaid_Electroclaw_Hatches":            "Люки под током",
	}
)

// OperatorStateTranslate return russian translate for skill
func OperatorStateTranslate(name string) string {
	name = strings.TrimPrefix(name, "operatorpvp_")
	idx := strings.Index(name, ":")
	if idx != -1 {
		name = name[:idx]
	}
	translate, ok := operatorSkillsTranslations[name]
	if ok {
		return translate
	}
	return name
}

type OperatorDefinition struct {
	ID              string `json:"id"`
	Category        string `json:"def"`
	Index           string `json:"index"`
	UniqueStatistic struct {
		PVP struct {
			StatisticID string `json:"statisticId"`
		} `json:"pvp"`
	} `json:"uniqueStatistic"`
}

type OperatorsDefinition struct {
	Definitions map[string]OperatorDefinition
	Indexes     map[string]string // index:name
}

func (od *OperatorsDefinition) UniqueStatsPVP() []string {
	result := make([]string, 0, len(od.Definitions))
	for _, definition := range od.Definitions {
		s := definition.UniqueStatistic.PVP.StatisticID
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

func (od *OperatorsDefinition) NameByID(idx string) string {
	name, ok := od.Indexes[idx]
	if ok {
		return name
	}
	return idx
}

func initOperatorsDefinition() error {
	Operators = &OperatorsDefinition{}
	err := json.Unmarshal([]byte(operatorJSON), &Operators.Definitions)
	if err != nil {
		return err
	}
	Operators.Indexes = make(map[string]string, len(Operators.Definitions))
	for name, definition := range Operators.Definitions {
		Operators.Indexes[definition.Index] = name
	}
	return nil
}

var (
	Operators *OperatorsDefinition
)

func init() {
	if err := initOperatorsDefinition(); err != nil {
		log.Println("Error caused by initOperatorsDefinition -", err)
	}
}
