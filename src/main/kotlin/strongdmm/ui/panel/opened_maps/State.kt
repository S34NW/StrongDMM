package strongdmm.ui.panel.opened_maps

import strongdmm.byond.dmm.Dmm
import strongdmm.service.action.ActionBalanceStorage

class State {
    lateinit var providedOpenedMaps: Set<Dmm>
    lateinit var providedActionBalanceStorage: ActionBalanceStorage

    var selectedMap: Dmm? = null
}
