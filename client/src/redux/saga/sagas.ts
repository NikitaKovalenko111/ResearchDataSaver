import { getResources, sendDocument, sendInternetArticle, sendLibraryArticle, sendFipsContent } from './../../api/resources-api';
import { call, put, takeLatest } from 'redux-saga/effects'
import type { Action } from '@reduxjs/toolkit';
import type { ItemPayload, ItemT, searchObj } from '../../types';
import { setResources } from '../features/resources/resourcesSlice';

interface GetResourcesAction extends Action<string> {
    type: 'resources/getResources';
    payload: {
        type: string;
        searchText: string
        searchObj: searchObj
    }
}

interface AddResourceAction extends Action<string> {
    type: 'resources/addResource';
    payload: ItemPayload
}

function* getResourcesSaga(action: GetResourcesAction) {
    // @ts-ignore
    const resources: ItemT[] = yield call(getResources, action.payload.type, action.payload.searchText, action.payload.searchObj);
    yield put(setResources(resources));
}

export function* getResourcesWatcher() {
    yield takeLatest('resources/getResources', getResourcesSaga);
}

function* addResourceSaga(action: AddResourceAction) {
    switch (action.payload.itemType) {
        case "document": {
            yield call(sendDocument, action.payload);
            break;
        }
        case "internet_article": {
            yield call(sendInternetArticle, action.payload);
            break;
        }
        case "library_article": {
            yield call(sendLibraryArticle, action.payload);
            break;
        }
        case "fips_content": {
            yield call(sendFipsContent, action.payload);
            break;
        }
    
        default:
            break;
    }
}

export function* addResourceSagaWatcher() {
    yield takeLatest('resources/addResource', addResourceSaga);
}