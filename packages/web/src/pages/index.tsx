import styled from '@emotion/styled';
import { PageWrapper } from 'ui/styles/components';
import Head from 'components/Head';
import Image from 'next/image';
import AboutWindow from 'features/about/components/AboutWindow';
import Cookies from 'js-cookie';
import DesktopIconList from 'components/DesktopIconList';
// import NewsWindow from 'features/news/components/NewsWindow';
// import { PostType } from 'features/news/types';
// import { getAllPosts } from 'utils/lib';

const IndexWrapper = styled(PageWrapper)`
  max-width: var(--content-width-xl);
`;

// const IndexPage = ({ allPosts }: { allPosts: PostType[] }) => {
const IndexPage = () => {
  return (
    <IndexWrapper>
      <Head />
      <Image src="/images/under-construction.gif" alt="Under Construction" layout="fill" />
    </IndexWrapper>
  );
};
export default IndexPage;
