import { DomainProp } from "../page"
import { IdeaInformation } from "../ui/idea-board/idea-post";

export const fetchDomains = async (): Promise<DomainProp[]> => {
    const res = await fetch(`http://localhost:8080/api/v1/domains/limit/14`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
    });
    return await res.json();
}

export const fetchDomainByID = async (domainID: string): Promise<DomainProp> => {
    const res = await fetch(`http://localhost:8080/api/v1/domains/${domainID}`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
    });
    return await res.json();
}

export const fetchIdeasByDomainID = async (domainID: string): Promise<IdeaInformation[]> => {
    const res = await fetch(`http://localhost:8080/api/v1/ideas/domain-id/${domainID}`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
    });
    return await res.json(); 
}
