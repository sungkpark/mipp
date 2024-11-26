import Link from "next/link";
import { DomainProp } from "../page";

export default function DomainBall({
    domainProp,
}: {
    domainProp: DomainProp;
}) {
    return (
        <Link
            href={`idea-board/${domainProp.domainId}`}
        >
            <div className="transition ease-in-out delay-50 rounded-full bg-[#FDE5BF] hover:bg-[#3A7B58] hover:-translate-y-1 hover:scale-110 duration-300">
                <div className="inline-block h-32 w-32 text-center align-middle leading-[8rem]">
                    {domainProp.domainName}
                </div>
            </div>
        </Link>
    );
}