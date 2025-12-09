import { useEffect, useState, type JSX } from "react";
import type React from "react";
import { Item } from "../../components/Item/Item";
import { Dialog } from "../../components/dialog/Dialog";
import { useDispatch, useSelector } from "react-redux";
import type { AppDispatch, RootState } from "../../redux/store";
import type { ItemType } from "../../types";

type PropsType = {}

export const ItemsPage: React.FC<PropsType> = (): JSX.Element => {
    const [searchText, setSearchText] = useState<string>("")
    const [isOpened, setIsOpened] = useState<boolean>(false)
    const [type, setType] = useState<ItemType>("library_article")

    const dispatch: AppDispatch = useDispatch()

    const addButtonClickHandler = () => {
        setIsOpened(true)
    }

    const resources = useSelector((state: RootState) => state.resources.resources)

    useEffect(() => {
        dispatch({
            type: 'resources/getResources',
            payload: {
                type: type,
                searchText: searchText
            }
        })
    }, [type, searchText])

    return (
        <>
            <Dialog isOpened={isOpened} setIsOpened={setIsOpened} />
            <div className="items">
                <div className="container">
                    <h1 className="items__header">Ресурсы исследования</h1>
                    <input onChange={(el) => {
                        setSearchText(el.target.value)
                    }} value={searchText} type="text" placeholder="Название ресурса" className="items__search-input" />
                    <select className="items__select" name="types" onChange={(e: React.ChangeEvent<HTMLSelectElement>) => {
                        setType(e.target.value as ItemType)
                    }} value={ type } id="items-types">
                        <option value="library_article">Научные статьи</option>
                        <option value="internet_article">Интернет статьи</option>
                        <option value="document">Документы</option>
                        <option value="fips_content">Патенты/программы для ЭВМ</option>
                    </select>
                    <div className="items__wrapper">
                        {
                            resources.length != 0 && resources.map(resource => (
                                <Item
                                    key={resource.id}
                                    id={resource.id}
                                    itemAnnotation={resource.itemAnnotation}
                                    itemLink={resource.itemLink}
                                    publishingDate={resource.publishingDate}
                                    itemAuthors={resource.itemAuthors}
                                    itemType={resource.itemType}
                                    itemName={resource.itemName}
                                    itemLang={resource.itemLang}
                                    itemAddress={resource.itemAddress}
                                    itemApplicant={resource.itemApplicant}
                                    itemFipsType={resource.itemFipsType}
                                    itemPublisher={resource.itemPublisher}
                                    itemPublisherObject={resource.itemPublisherObject}
                                    itemRegistration={resource.itemRegistration}
                                    itemSearchingMachine={resource.itemSearchingMachine}
                                    itemSupervisor={resource.itemSupervisor}
                                    itemUDK={resource.itemUDK}
                                />
                            ))
                        }
                    </div>
                    <button onClick={addButtonClickHandler} className="items__add-button">Добавить</button>
                </div>
            </div>
        </>
    )
}