import { Button } from "../button";

export default function IdeaPostWrapper() {
    return (
        <div className="flex flex-col">
            <div className="justify-items-stretch">
                <IdeaPost title="asdfasdfa" description="adsjfald"></IdeaPost>
            </div>
            <div className="justify-items-center p-6">
                <Button>Vote</Button>
            </div>
        </div>
    );
}

export function IdeaPost({
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