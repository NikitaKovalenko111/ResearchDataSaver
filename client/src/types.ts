export type ItemType = "document" | "internet_article" | "library_article" | "fips_content"

export type ItemT = {
    id: number
    itemAnnotation: string
    itemLink: string
    publishingDate: string
    itemLang?: string
    itemUDK?: string
    itemPublisherObject?: string
    itemPublisher?: string
    itemAuthors: string[]
    itemSearchingMachine?: string
    itemSupervisor?: string
    itemApplicant?: string
    itemAddress?: string
    itemFipsType?: string
    itemRegistration?: string
    itemType: ItemType
    itemName: string
}

export interface searchObj {
    fipsType?: 'patent' | 'program' | 'all'
    lang?: string
    date?: string
    supervisor?: string
    reg?: string
    sm?: string
}

export type ItemPayload = {
    itemAnnotation: string
    itemLink: string
    publishingDate: string
    itemLang?: string
    itemUDK?: string
    itemPublisherObject?: string
    itemPublisher?: string
    itemAuthors: string[]
    itemSearchingMachine?: string
    itemSupervisor?: string
    itemApplicant?: string
    itemAddress?: string
    itemFipsType?: string
    itemRegistration?: string
    itemType: ItemType
    itemName: string
}