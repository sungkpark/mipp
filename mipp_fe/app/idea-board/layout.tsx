import CompanyInformation from "../ui/idea-board/company-information"
import MippitStatus from "../ui/idea-board/mippit-status"
import VotedIdeads from "../ui/idea-board/voted-ideas"

export default function Layout({ children }: { children: React.ReactNode }) {
    return (
        <div>
            {children}           
        </div>
    )
}