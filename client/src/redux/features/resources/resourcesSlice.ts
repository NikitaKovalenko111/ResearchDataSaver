import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import type { ItemT } from '../../../types'

export interface ResourcesState {
    resources: ItemT[]
}

const initialState: ResourcesState = {
  resources: [],
}

export const resourcesSlice = createSlice({
  name: 'resources',
  initialState,
  reducers: {
    setResources: (state, action: PayloadAction<ItemT[]>) => {
        state.resources = action.payload
    }
  },
})

export const { setResources } = resourcesSlice.actions

export default resourcesSlice.reducer