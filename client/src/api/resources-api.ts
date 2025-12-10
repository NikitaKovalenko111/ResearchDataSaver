import axios from "axios";
import type { ItemPayload, ItemT, ItemType, searchObj } from "../types";

const instance = axios.create({
    baseURL: "http://localhost:3001",
    headers: {
        "Content-Type": "application/json",
    },
})

export const getResources = async (type: ItemType, searchText: string, searchObj: searchObj): Promise<ItemT[]> => {
    const { fipsType, lang, reg, supervisor, date, sm } = searchObj

    console.log(searchObj);
    

    switch (type) {
        case "document": {
            const response = await instance.get(`/document/all?name=${searchText}${(date != '') ? `&date=${date}` : ''}`);

            const responseObj: ItemT[] = response.data ? [
                ...response.data.map((el: any): ItemT => {
                    return {
                        id: el.id,
                        itemAnnotation: el.annotation,
                        publishingDate: el.publishingDate,
                        itemLink: el.link,
                        itemType: 'document',
                        itemName: el.name,
                        itemAuthors: [el.author]
                    }
                })
            ] : []

            return responseObj
        }
        case "fips_content": {
            const response = await instance.get(`/fips/all?name=${searchText}${(fipsType == 'patent' || fipsType == 'program') ? `&fipsType=${fipsType}` : ''}${(reg != '') ? `&reg=${reg}` : ''}${(date != '') ? `&date=${date}` : ''}`);

            const responseObj: ItemT[] = response.data ? [
                ...response.data.map((el: any): ItemT => {
                    return {
                        id: el.id,
                        itemAnnotation: el.annotation,
                        publishingDate: el.publishingDate,
                        itemLink: el.link,
                        itemType: 'fips_content',
                        itemName: el.name,
                        itemAuthors: el.authors,
                        itemApplicant: el.applicant,
                        itemAddress: el.address,
                        itemRegistration: el.registration,
                        itemFipsType: el.type
                    }
                })
            ] : []
            
            return responseObj
        }
        case "internet_article": {
            const response = await instance.get(`/internet-article/all?name=${searchText}${(date != '') ? `&date=${date}` : ''}${(sm != '') ? `&sm=${sm}` : ''}`);
            const responseObj: ItemT[] = response.data ? [
                ...response.data.map((el: any): ItemT => {
                    return {
                        id: el.id,
                        itemAnnotation: el.annotation,
                        itemSearchingMachine: el.searchingMachine,
                        publishingDate: el.publishingDate,
                        itemLink: el.link,
                        itemType: 'internet_article',
                        itemName: el.name,
                        itemAuthors: [el.author]
                    }
                })
            ] : []
            
            return responseObj
        }
        case "library_article": {
            const response = await instance.get(`/library-article/all?name=${searchText}${(lang != '') ? `&lang=${lang}` : ''}${(supervisor != '') ? `&supervisor=${supervisor}` : ''}${(date != '') ? `&date=${date}` : ''}`);

            const responseObj: ItemT[] = response.data ? [
                ...response.data.map((el: any): ItemT => {
                    return {
                        id: el.id,
                        itemAnnotation: el.annotation,
                        publishingDate: el.publishingDate,
                        itemLink: el.link,
                        itemName: el.name,
                        itemType: 'library_article',
                        itemAuthors: el.authors,
                        itemLang: el.lang,
                        itemUDK: el.udk,
                        itemPublisherObject: el.publisherObject,
                        itemPublisher: el.publisher,
                        itemSupervisor: el.supervisor
                    }
                })
            ] : []
            
            return responseObj
        }
    
        default:
            return [];
    }
}

export const sendInternetArticle = async (data: ItemPayload): Promise<void> => {
    await instance.post("/internet-article/create", {
        name: data.itemName,
        annotation: data.itemAnnotation,
        link: data.itemLink,
        publishingDate: data.publishingDate,
        author: data.itemAuthors[0],
        searchingMachine: data.itemSearchingMachine
    });
}

export const sendLibraryArticle = async (data: ItemPayload): Promise<void> => {
    await instance.post("/library-article/create", {
        name: data.itemName,
        annotation: data.itemAnnotation,
        link: data.itemLink,
        publishingDate: data.publishingDate,
        authors: data.itemAuthors,
        udk: data.itemUDK,
        lang: data.itemLang,
        publisherObject: data.itemPublisherObject,
        publisher: data.itemPublisher,
        supervisor: data.itemSupervisor,
    });
}

export const sendFipsContent = async (data: ItemPayload): Promise<void> => {
    await instance.post("/fips/create", {
        name: data.itemName,
        annotation: data.itemAnnotation,
        link: data.itemLink,
        publishingDate: data.publishingDate,
        authors: data.itemAuthors,
        registration: data.itemRegistration,
        type: data.itemFipsType,
        applicant: data.itemApplicant,
        address: data.itemAddress,
    });
}

export const sendDocument = async (data: ItemPayload): Promise<void> => {
    await instance.post("/document/create", {
        name: data.itemName,
        annotation: data.itemAnnotation,
        link: data.itemLink,
        publishingDate: data.publishingDate,
        author: data.itemAuthors[0],
    });
}