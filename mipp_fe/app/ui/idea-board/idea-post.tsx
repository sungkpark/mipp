import { Button } from "../button";

export interface IdeaInformation {
    ideaID: number;
    title: string;
    description: string;
    capturedUrl: string;
}

export default function IdeaPost({
    title,
    description,
}: {
    title: string,
    description: string,
}) {
    return (
        <div className="bg-[#FDE5BF] rounded-2xl h-96">
            <div className="p-6">
                <div className="font-bold">
                    {title}
                </div>
                <div className="line-clamp-6">
                    {description}
                </div>
            </div>
        </div>
    );
}