// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import RarimoRarimoCoreRarimoRarimocoreRarimocore from './rarimo/rarimo-core/rarimo.rarimocore.rarimocore'
import RarimoRarimoCoreRarimoRarimocoreTokenmanager from './rarimo/rarimo-core/rarimo.rarimocore.tokenmanager'


export default { 
  RarimoRarimoCoreRarimoRarimocoreRarimocore: load(RarimoRarimoCoreRarimoRarimocoreRarimocore, 'rarimo.rarimocore.rarimocore'),
  RarimoRarimoCoreRarimoRarimocoreTokenmanager: load(RarimoRarimoCoreRarimoRarimocoreTokenmanager, 'rarimo.rarimocore.tokenmanager'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
