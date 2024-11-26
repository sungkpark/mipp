import IdeaPostWrapper from "../../ui/idea-board/idea-post";
import { LuRefreshCw } from "react-icons/lu";
import { fetchDomainByID, fetchIdeasByDomainID } from "@/app/lib/api";
import CompanyInformation from "@/app/ui/idea-board/company-information";
import MippitStatus from "@/app/ui/idea-board/mippit-status";
import VotedIdeads from "@/app/ui/idea-board/voted-ideas";
import IdeaSelection from "@/app/ui/idea-board/idea-selection";

export default async function Page({
    params,
}: {
    params: { slug: string }
}) {
    const slugDomainID = (await params).slug;

    let domainProp = await fetchDomainByID(slugDomainID);

    return (
        <div className="flex flex-row h-screen md:overflow-hiddne">
            <div className="flex-auto basis-1/5 mx-8 my-4">
                <CompanyInformation domainID={domainProp.domainId} information={domainProp.companyInformation}/>
            </div>

            <div className="flex-auto basis-3/5">
                <IdeaSelection domainProp={domainProp}></IdeaSelection>
            </div>

            <div className="flex-auto basis-1/5 mx-8 my-4">
                <div className="flex flex-col gap-3">
                    <MippitStatus domainID={domainProp.domainId}/>
                    <VotedIdeads domainID={domainProp.domainId}/>
                </div>
            </div>
        </div>
    );
}