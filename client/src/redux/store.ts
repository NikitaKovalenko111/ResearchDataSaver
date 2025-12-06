import { configureStore } from '@reduxjs/toolkit'
import { resourcesSlice } from './features/resources/resourcesSlice'
import createSagaMiddleware from 'redux-saga';
import { addResourceSagaWatcher, getResourcesWatcher } from './saga/sagas';

const sagaMiddleware = createSagaMiddleware();

export const store = configureStore({
  reducer: {
    resources: resourcesSlice.reducer,
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(sagaMiddleware),
})

sagaMiddleware.run(getResourcesWatcher);
sagaMiddleware.run(addResourceSagaWatcher);

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch