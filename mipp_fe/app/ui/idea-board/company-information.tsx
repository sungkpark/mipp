import { CheckBadgeIcon, CheckIcon } from "@heroicons/react/24/outline";

export default function CompanyInformation({
    domainID,
    information,
}: {
    domainID: number;
    information: string;
}) {
    return (
        <div className="h-4/5 bg-[#FDE5BF] rounded-2xl">
            <div className="p-6">
                <div className="text-center font-bold">Company information</div>
                <div>
                    <p>{information}</p>
                </div>
                <div className="text-center font-bold">Company activity</div>
                <div>
                    <div className="flex items-center justify-between">
                        <div>
                            Has admin/verified
                        </div>
                        <div className="size-5">
                            <CheckBadgeIcon></CheckBadgeIcon>
                        </div>
                    </div>
                    <div className="flex items-center justify-between">
                        <div>
                            Active in the last 30 days
                        </div>
                        <div className="size-5">
                            <CheckIcon></CheckIcon>
                        </div>
                    </div>
                    <div className="flex items-center justify-between">
                        <div>
                            Number of ideas selected
                        </div>
                        <div>
                            6
                        </div>
                    </div>
                    <div className="flex items-center justify-between">
                        <div>
                            Number of ideas in progress
                        </div>
                        <div>
                            3
                        </div>
                    </div>
                    <div className="flex items-center justify-between">
                        <div>
                            Number of ideas completed
                        </div>
                        <div>
                            2
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}