import { mali } from '@/app/ui/font/fonts';
import { fetchDomains } from './lib/api';
import DomainBall from './ui/domain-balls';

export interface DomainProp {
  domainId: number;
  domainName: string;
  companyInformation: string;
  verified: boolean;
}

export default async function Page() {
  let domainPageProps = await fetchDomains()

  return (
    <main className="flex items-center min-h-screen flex-col p-8">
      <div className="flex flex-col gap-4 w-3/4 md:flex-row">
        <div className='m-12 flex flex-col justfify-center rounded-lg px-6 py-10'>
          <p className={`${mali.className} text-xl md:text-3xl`}>
            Welcome to Mippit. Mippit is a platform that helps both end of the user journey on the web. By posting an feature idea using the chrome extension, we put closer the reach of innovation on the website domain.
            <br/><br/>
            Here are some websites with feature ideas posted. Click and check out what other people have posted. Or, search for a specific company on the top right.
          </p>
        </div>
      </div>
      <div className="grid grid-rows-2 grid-flow-col gap-12">
        {domainPageProps.map((domainProp) => (
          <DomainBall domainProp={domainProp}></DomainBall>
        ))}
      </div>
    </main>
  );
}
