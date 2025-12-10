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
    const [sm, setSM] = useState<string>("")
    const [lang, setLang] = useState<string>("")
    const [supervisor, setSupervisor] = useState<string>("")
    const [date, setDate] = useState<string>("")
    const [reg, setReg] = useState<string>("")
    const [fipsType, setFipsType] = useState<'program' | 'patent' | 'all'>('all')

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
                searchText: searchText,
                searchObj: {
                    fipsType: fipsType,
                    lang: lang,
                    date: date,
                    supervisor: supervisor,
                    reg: reg,
                    sm: sm,
                }
            }
        })
    }, [type, searchText, sm, lang, reg, fipsType, supervisor, date])

    return (
        <>
            <Dialog isOpened={isOpened} setIsOpened={setIsOpened} />
            <div className="items">
                <div className="container">
                    <h1 className="items__header">Ресурсы исследования</h1>
                    <input onChange={(el) => {
                        setSearchText(el.target.value)
                    }} value={searchText} type="text" placeholder="Название ресурса" className="items__search-input" />
                    {
                        type == "internet_article" && (
                            <div className="items__input-wrapper">
                                <input onChange={(el) => {
                                    setSM(el.target.value)
                                }} value={sm} type="text" placeholder="Поисковая машина" className="items__search-input" />
                                <input onChange={(el) => {
                                setDate(el.target.value)
                                }} value={date} type="date" placeholder="Дата публикации" className="items__search-input" />
                            </div>
                        )
                    }
                    {
                        type == 'library_article' && (
                            <div className="items__input-wrapper">
                                <input onChange={(el) => {
                                setLang(el.target.value)
                                }} value={lang} type="text" placeholder="Язык" className="items__search-input" />
                                <input onChange={(el) => {
                                setSupervisor(el.target.value)
                                }} value={supervisor} type="text" placeholder="Научный руководитель" className="items__search-input" />
                                <input onChange={(el) => {
                                setDate(el.target.value)
                                }} value={date} type="date" placeholder="Дата публикации" className="items__search-input" />
                            </div>
                        )
                    }
                    {
                        type == 'fips_content' && (
                            <div className="items__input-wrapper">
                                <input onChange={(el) => {
                                setReg(el.target.value)
                                }} value={reg} type="text" placeholder="Регистрация" className="items__search-input" />
                                <select className="items__search-input" name="fips-type" onChange={(el) => {
                                    setFipsType(el.target.value as ('all' | 'patent' | 'program'))
                                }} value={fipsType} id="fips-type">
                                    <option value="program" className="dialog__form-type-option">Программа для ЭВМ</option>
                                    <option value="patent" className="dialog__form-type-option">Патент</option>
                                    <option value="all" className="dialog__form-type-option">Все</option>
                                </select>
                                <input onChange={(el) => {
                                setDate(el.target.value)
                                }} value={date} type="date" placeholder="Дата публикации" className="items__search-input" />
                            </div>
                        )
                    }
                    {
                        type == 'document' && (
                            <input onChange={(el) => {
                                setDate(el.target.value)
                                }} value={date} type="date" placeholder="Дата публикации" className="items__search-input" />
                        )
                    }
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