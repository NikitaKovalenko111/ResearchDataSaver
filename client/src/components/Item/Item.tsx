import type { JSX } from "react";
import type React from "react";
import { type ItemT } from "../../types";

type PropsType = ItemT

export const Item: React.FC<PropsType> = ({ itemName, itemType, publishingDate, itemAnnotation, itemAuthors, itemLink, itemUDK, itemPublisherObject, itemSupervisor, itemLang, itemRegistration, itemAddress, itemApplicant, itemFipsType, itemSearchingMachine }): JSX.Element => {
    return (
        <div className="item">
            <div className="item__data">
                {
                    itemType != "fips_content" && (
                        <span className="item__type item__property">{ itemType == 'document' ? "Документ" : itemType == 'internet_article' ? "Интернет статья" : itemType == 'library_article' ? "Научная статья" : "" }</span>
                    )
                }
                {
                    itemType == "library_article" && (
                        <>
                            <span className="item__property">УДК: <span className="item__span-bold">{ itemUDK }</span></span>
                            <span className="item__property">Язык: <span className="item__span-bold">{ itemLang }</span></span>
                            <span className="item__property">Науч. рук: <span className="item__span-bold">{ itemSupervisor }</span></span>
                            <span className="item__property"><span className="item__span-bold">{ itemPublisherObject }</span></span>
                        </>
                    )
                }
                {
                    itemType == "fips_content" && (
                        <>
                            <span className="item__property">Регистрация: <span className="item__span-bold">{ itemRegistration }</span></span>
                            <span className="item__property">Адрес: <span className="item__span-bold">{ itemAddress }</span></span>
                            <span className="item__property"><span className="item__span-bold">{ itemApplicant }</span></span>
                        </>
                    )
                }
                {
                    itemType == "internet_article" && (
                        <span className="item__property">Поисковая машина: <span className="item__span-bold">{ itemSearchingMachine }</span></span>
                    )
                }
                <span className="item_publishing-date item__property">Дата публикации: <span className="item__span-bold">{ (new Date(publishingDate)).toLocaleDateString() }</span></span>
            </div>
            <span className="item__name">Название: { itemName }</span>
            {
                itemType == "fips_content" && (
                    <span className="item__property">Тип: <span className="item__span-bold">{ itemFipsType == 'patent' ? 'Патент' : "Программа для ЭВМ" }</span></span>
                )
            }
            <div className="item__authors">
            {
                itemAuthors.map(el => {
                    return <span className="item_property"><span className="item__span-bold">{ el }</span></span>
                })
            }
            </div>
            <p className="item_annotation">
                { itemAnnotation }
            </p>
            <a href={ itemLink } className="item__open-button">Открыть</a>
        </div>
    )
}