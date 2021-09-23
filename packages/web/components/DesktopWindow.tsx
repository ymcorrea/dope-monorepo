import { media } from '../styles/mixins';
import ConditionalWrapper from './ConditionalWrapper';
import DesktopWindowTitleBar from './DesktopWindowTitleBar';
import Draggable from 'react-draggable';
import React from 'react';
import styled from '@emotion/styled';

interface WindowProps {
  width?: number;
  height?: number;
  title: string | undefined;
  titleChildren?: React.ReactNode;
  children: React.ReactNode;
}

const DesktopWindow = ({
  title,
  width = 1024,
  height = 768,
  titleChildren,
  children,
}: WindowProps) => {
  const WindowWrapper = styled.div`
    width: 100%;
    height: 100%;
    margin: 0;
    max-width: ${width}px;
    max-height: ${height}px;
    ${media.tablet`
        width: 100%;
        height: 100%;
        margin: 0;
      `}
    ${media.laptop`
        width: 80%;
        height: 90%;
        margin: auto;
        margin-top: 32px;
      `}
    padding: 0;
    background-color: #a8a9ae;
    border: 2px solid #000;
    filter: drop-shadow(8px 8px rgba(0, 0, 0, 0.15));
    display: flex;
    flex-direction: column;
  `;

  const isTouchDevice = () => {
    if (typeof window === 'undefined') {
      return false;
    }
    return (
      'ontouchstart' in window || navigator.maxTouchPoints > 0 || navigator.msMaxTouchPoints > 0
    );
  };

  return (
    <ConditionalWrapper
      condition={!isTouchDevice()}
      wrap={children => <Draggable handle=".windowTitleBar">{children}</Draggable>}
    >
      <WindowWrapper>
        <DesktopWindowTitleBar title={title}>{titleChildren}</DesktopWindowTitleBar>
        {children}
      </WindowWrapper>
    </ConditionalWrapper>
  );
};
export default DesktopWindow;