import CompanyInformation from "../ui/idea-board/company-information"
import MippitStatus from "../ui/idea-board/mippit-status"
import VotedIdeads from "../ui/idea-board/voted-ideas"

export default function Layout({ children }: { children: React.ReactNode }) {
    return (
        <div className="flex flex-row h-screen md:overflow-hiddne">
            <div className="flex-auto basis-1/5 mx-8 my-4">
                <CompanyInformation information=""/>
            </div>
            <div className="flex-auto basis-3/5">
                {children}
            </div>
            <div className="flex-auto basis-1/5 mx-8 my-4">
                <div className="flex flex-col gap-3">
                    <MippitStatus/>
                    <VotedIdeads/>
                </div>
            </div>
        </div>
    )
}