package main

var samples = `
{
    "contractState": {
        "activeAssets": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "nickname": "HeavyEquip",
        "version": "The version number of the current contract"
    },
    "event": {
        "AssetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "carrier": "transport entity currently in possession of asset",
        "runhours": "number of hours equipment used",
        "temperature": "Temperature of the asset in CELSIUS.",
        "timestamp": "2017-04-11T10:12:03.082983911+05:30"
    },
    "initEvent": {
        "nickname": "HeavyEquip",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "state": {
        "AssetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "alerts": {
            "active": [
                "OVERTTEMP"
            ],
            "cleared": [
                "OVERTTEMP"
            ],
            "raised": [
                "OVERTTEMP"
            ]
        },
        "carrier": "transport entity currently in possession of asset",
        "inCompliance": true,
        "lastEvent": {
            "args": [
                "parameters to the function, usually args[0] is populated with a JSON encoded event object"
            ],
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
        },
        "runhours": 123.456,
        "temperature": 123.456,
        "timestamp": "2017-04-11T10:12:03.0841114+05:30"
    }
}`