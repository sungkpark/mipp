import IdeaPostWrapper from "../../ui/idea-board/idea-post";
import { LuRefreshCw } from "react-icons/lu";

export default async function Page({
    params,
}: {
    params: Promise<{ slug: string }>
}) {
    const slug = (await params).slug

    return (
        <div>
            <div className="text-center text-4xl my-4 font-bold">
                {slug}
            </div>
            <div className="flex flex-row-reverse items-center gap-2 p-3">
                <p>refresh ideas</p>
                <LuRefreshCw/>
            </div>
            <div className="flex gap-4">
                <div className="flex-auto">
                    <IdeaPostWrapper/>
                </div>
                <div className="flex-auto">
                    <IdeaPostWrapper/>
                </div>
                <div className="flex-auto">
                    <IdeaPostWrapper/>
                </div>
            </div>
        </div>
    );
}