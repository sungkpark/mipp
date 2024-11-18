import '@/app/ui/global.css';
import { mali } from '@/app/ui/font/fonts';
import Search from './ui/search';
import MippitLogo from './ui/mippit-logo';
import Link from 'next/link';

export default function RootLayout({ children, }: { children: React.ReactNode; }) {
  return (
    <html lang="en">
      <body className={`${mali.className} antialiased bg-[#F8B76C]`}>
        <div className='flex flex-row'>
          <div className='basis-1/12 md:basis-6/12'></div>
          <Link
            href='/'
            className='basis-1/12 my-5 p-3'
            // className='basis-1/12 my-5 rounded-2xl bg-[#3A7B58] p-3'
          >
            <MippitLogo/>
          </Link>
          <div className='basis-3/12'></div>
          <div className={`${mali.className} basis-3/12 md:basis-2/12 my-5`}>
            <Search placeholder='Search for domain'/>
          </div>
          <div className='basis-1/12'></div>
        </div>
        {children}
      </body>
    </html>
  );
}
