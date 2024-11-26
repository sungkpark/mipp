import { FaComment } from "react-icons/fa";

export default function VotedIdeads({domainID}: {
    domainID: number;
}) {
    return (
        <div className="h-3/5 bg-[#FDE5BF] rounded-2xl">
            <div className="p-6">
                <div className="text-center font-bold">To the threads</div>
                <div className="flex items-center justify-between">
                    <div className="line-clamp-1">Title of the idea you've voted</div>
                    <div className="flex items-center justify-between gap-1">
                        <FaComment/><p>32</p>
                    </div>
                </div>
                <div className="flex items-center justify-between">
                    <div className="line-clamp-1">Title of the idea you've voted</div>
                    <div className="flex items-center justify-between gap-1">
                        <FaComment/><p>12</p>
                    </div>
                </div>
                <div className="flex items-center justify-between">
                    <div className="line-clamp-1">Title of the idea you've voted and this title is long to see line clamp</div>
                    <div className="flex items-center justify-between gap-1">
                        <FaComment/><p>3</p>
                    </div>
                </div>
            </div>
        </div>
    );
}