import { useState, type Dispatch, type JSX, type SetStateAction } from "react";
import type React from "react";
import cn from 'classnames'
import { useDispatch } from "react-redux";
import type { ItemPayload, ItemType } from "../../types";

type PropsType = {
    isOpened: boolean
    setIsOpened: Dispatch<SetStateAction<boolean>>
}

export const Dialog: React.FC<PropsType> = ({ isOpened, setIsOpened }): JSX.Element => {
    const [resourceType, setResourceType] = useState<ItemType>("internet_article")
    const [resourceName, setResourceName] = useState<string>("")
    const [resourceLink, setResourceLink] = useState<string>("")
    const [resourceAnnotation, setResourceAnnotation] = useState<string>("")
    const [resourceSM, setResourceSM] = useState<string>("")
    const [resourceLang, setResourceLang] = useState<string>("")
    const [resourceUDK, setResourceUDK] = useState<string>("")
    const [publisherObject, setPublisherObject] = useState<string>("")
    const [publisher, setPublisher] = useState<string>("")
    const [supervisor, setSupervisor] = useState<string>("")
    const [fipsType, setFipsType] = useState<"program" | "patent">("program")
    const [registration, setRegistration] = useState<string>("")
    const [applicant, setApplicant] = useState<string>("")
    const [address, setAddress] = useState<string>("")
    const [resourceDate, setResourceDate] = useState<string>("")
    const [authorInputs, setAuthorInputs] = useState<string[]>([])

    const dispatch = useDispatch()

    const changeResourceDateHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setResourceDate(e.target.value)
    }

    const changeResourceTypeHandler = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setResourceType(e.target.value as ItemType)
    }

    const changeAddressHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setAddress(e.target.value)
    }

    const changeResourceNameHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setResourceName(e.target.value)
    }

    const changeApplicantHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setApplicant(e.target.value)
    }

    const changeRegistrationHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setRegistration(e.target.value)
    }

    const changeFipsTypeHandler = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setFipsType(e.target.value as ("program" | "patent"))
    }

    const changeResourceLangHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setResourceLang(e.target.value)
    }

    const changeResourceUDKHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setResourceUDK(e.target.value)
    }

    const changePublisherObjectHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPublisherObject(e.target.value)
    }

    const changePublisherHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPublisher(e.target.value)
    }

    const changeSupervisorHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSupervisor(e.target.value)
    }

    const changeResourceLinkHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setResourceLink(e.target.value)
    }

    const changeResourceAnnotationHandler = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        setResourceAnnotation(e.target.value)
    }

    const changeResourceSMHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        setResourceSM(e.target.value)
    }

    const handleClose = (e: React.MouseEvent) => {
        if (e.target == e.currentTarget) {
            setIsOpened(false)
        }
    }

    const submitHandler = (e: React.FormEvent) => {
        e.preventDefault()

        const data: ItemPayload = {
            itemAnnotation: resourceAnnotation,
            itemLink: resourceLink,
            publishingDate: resourceDate,
            itemLang: resourceLang,
            itemUDK: resourceUDK,
            itemPublisherObject: publisherObject,
            itemPublisher: publisher,
            itemAuthors: authorInputs,
            itemSearchingMachine: resourceSM,
            itemSupervisor: supervisor,
            itemApplicant: applicant,
            itemAddress: address,
            itemFipsType: fipsType,
            itemRegistration: registration,
            itemType: resourceType,
            itemName: resourceName
        }

        dispatch({
            type: 'resources/addResource',
            payload: data
        })

        setIsOpened(false)
    }

    return (
        <div onClick={ handleClose } className={cn('dialog', { 'dialog__active': isOpened })}>
            <div className="dialog__wrapper">
                <h2 className="dialog__header">Добавить ресурс</h2>

                <form className="dialog__form" method="post">
                    <label htmlFor="#">
                        <span className="dialog__form-label">Тип:</span>
                        <select name="type" onChange={ changeResourceTypeHandler } value={resourceType} id="type">
                            <option value="internet_article" className="dialog__form-type-option">Интернет статья</option>
                            <option value="library_article" className="dialog__form-type-option">Научная статья</option>
                            <option value="fips_content" className="dialog__form-type-option">Патент/программа для ЭВМ</option>
                            <option value="document" className="dialog__form-type-option">Правовой документ</option>
                        </select>
                    </label>
                    <input type="text" onChange={ changeResourceNameHandler } value={ resourceName } placeholder="Название" name="resource-name" className="dialog__form-input" />
                    <input type="date" onChange={ changeResourceDateHandler } value={ resourceDate } name="resource-date" id="date" />
                    <div className="dialog__authors">
                        <label htmlFor="authors">
                            <span className="dialog__form-label">Авторы:</span>
                        </label>
                        <div className="dialog__author-inputs">
                            {authorInputs?.map((author, index) => (
                                <input
                                    key={index}
                                    type="text"
                                    placeholder={`Автор ${index + 1}`}
                                    value={author}
                                    onChange={(e) => {
                                        const newAuthors = [...authorInputs];
                                        newAuthors[index] = e.target.value;
                                        setAuthorInputs(newAuthors);
                                    }}
                                    name={`author-${index}`}
                                    className="dialog__form-input"
                                />
                            ))}
                        </div>
                        {authorInputs.length < 10 && (
                            <button
                                type="button"
                                onClick={() => setAuthorInputs([...(authorInputs || []), ""])}
                                className="dialog__form-button-add"
                            >
                                + Добавить автора
                            </button>
                        )}
                    </div>
                    {
                        resourceType == "internet_article" && (
                            <>
                            <input type="text" onChange={ changeResourceSMHandler } value={ resourceSM } placeholder="Поисковая машина" name="resource-sm" className="dialog__form-input" />
                            </>
                        )
                    }
                    {
                        resourceType == "library_article" && (
                            <>
                            <div className="dialog__additional">
                                <input type="text" onChange={ changeResourceLangHandler } value={ resourceLang } placeholder="Язык статьи" name="resource-lang" className="dialog__form-input" />
                                <input type="text" onChange={ changeResourceUDKHandler } value={ resourceUDK } placeholder="УДК статьи" name="resource-udk" className="dialog__form-input" />
                                <input type="text" onChange={ changePublisherObjectHandler } value={ publisherObject } placeholder="Учебное заведение" name="publisher-object" className="dialog__form-input" />
                                <input type="text" onChange={ changePublisherHandler } value={ publisher } placeholder="Издательство" name="publisher" className="dialog__form-input" />
                                <input type="text" onChange={ changeSupervisorHandler } value={ supervisor } placeholder="Научный руководитель" name="supervisor" className="dialog__form-input" />
                            </div>
                            </>

                        )
                    }
                    {
                        resourceType == "fips_content" && (
                            <div className="dialog__additional">
                                <select name="fips-type" onChange={ changeFipsTypeHandler } value={fipsType} id="fips-type">
                                    <option value="program" className="dialog__form-type-option">Программа для ЭВМ</option>
                                    <option value="patent" className="dialog__form-type-option">Патент</option>
                                </select>
                               <input type="text" onChange={ changeRegistrationHandler } value={ registration } placeholder="Номер регистрации" name="registration" className="dialog__form-input" />
                               <input type="text" onChange={ changeApplicantHandler } value={ applicant } placeholder="Заявитель" name="applicant" className="dialog__form-input" />
                                <input type="text" onChange={ changeAddressHandler } value={ address } placeholder="Адрес" name="address" className="dialog__form-input" />
                            </div>
                        )
                    }
                    <input type="text" onChange={ changeResourceLinkHandler } value={ resourceLink } placeholder="Ссылка" name="resource-link" className="dialog__form-input" />
                    <textarea className="dialog__form-annotation" onChange={ changeResourceAnnotationHandler } value={ resourceAnnotation } name="resource-annotation" placeholder="Аннотация" id="resource-annotation"></textarea>

                    <button onClick={ submitHandler } className="dialog__form-submit">Добавить</button>
                </form>
            </div>
        </div>
    )
}