'use client'

import { DomainProp } from "@/app/page"
import { useEffect, useState } from "react";
import { LuRefreshCw } from "react-icons/lu";
import { IdeaInformation } from "./idea-post";
import { fetchIdeasByDomainID } from "@/app/lib/api";
import IdeaPost from "./idea-post";
import { Button } from "../button";

export default function IdeaSelection({domainProp}: {
    domainProp: DomainProp;
}) {
    const [ideaPageProps, setIdeaPageProps] = useState<IdeaInformation[]>([]);

    const refreshIdeas = async () => {
        const refreshedIdeas = await fetchIdeasByDomainID(domainProp.domainId.toString());
        setIdeaPageProps(refreshedIdeas);
    }

    useEffect(() => {
        refreshIdeas();
    }, [])

    return (
        <div>
            <div className="flex-auto basis-3/5">
                <div className="text-center text-4xl my-4 font-bold">
                    {domainProp?.domainName}
                </div>
                <div className="flex flex-row-reverse items-center gap-2 p-3 hover:cursor-pointer" onClick={refreshIdeas}>
                    <p>refresh ideas</p>
                    <LuRefreshCw />
                </div>
                <div className="flex gap-4">
                    {ideaPageProps.map((ideaPageProp, index) => (
                        <div key={index} className="flex-auto">
                            <div className="flex flex-col">
                                <div className="justify-items-stretch">
                                <IdeaPost title={ideaPageProp.title} description={ideaPageProp.description}></IdeaPost>
                            </div>
                            <div className="justify-items-center p-6">
                                <Button>Vote</Button>
                            </div>
                        </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )
}