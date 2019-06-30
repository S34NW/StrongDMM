package io.github.spair.strongdmm.logic.action

class MultipleAction(private val actions: List<Undoable>) : Undoable {
    override fun doAction(): Undoable {
        val reverseActions = mutableListOf<Undoable>()
        actions.forEach { reverseActions.add(it.doAction()) }
        return MultipleAction(reverseActions)
    }
}