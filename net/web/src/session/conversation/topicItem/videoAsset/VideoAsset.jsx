import { Modal } from 'antd';
import ReactResizeDetector from 'react-resize-detector';
import { VideoCameraOutlined } from '@ant-design/icons';
import { VideoAssetWrapper } from './VideoAsset.styled';
import { useVideoAsset } from './useVideoAsset.hook';

export function VideoAsset({ thumbUrl, lqUrl, hdUrl }) {

  const { state, actions } = useVideoAsset();

  const activate = () => {
    if (state.dimension.width / state.dimension.height > window.innerWidth / window.innerHeight) {
      let width = Math.floor(window.innerWidth * 8 / 10);
      let height = Math.floor(width * state.dimension.height / state.dimension.width);
      actions.setActive(width, height);
    }
    else {
      let height = Math.floor(window.innerHeight * 8 / 10);
      let width = Math.floor(height * state.dimension.width / state.dimension.height);
      actions.setActive(width, height);
    }
  }

  return (
    <VideoAssetWrapper>
      <ReactResizeDetector handleWidth={true} handleHeight={true}>
        {({ width, height }) => {
          if (width !== state.dimension.width || height !== state.dimension.height) {
            actions.setDimension({ width, height });
          }
          return <img style={{ height: '100%', objectFit: 'contain' }} src={thumbUrl} alt="" />
        }}
      </ReactResizeDetector>
      <div class="overlay" style={{ width: state.dimension.width, height: state.dimension.height }}>
        { !state.active && (
          <div onClick={activate}>
            <VideoCameraOutlined style={{ fontSize: 32, color: '#eeeeee', cursor: 'pointer' }} />
          </div>
        )}
        <Modal centered={true} style={{ backgroundColor: '#aacc00', padding: 0 }} visible={state.active} width={state.width + 12} bodyStyle={{ paddingBottom: 0, paddingTop: 6, paddingLeft: 6, paddingRight: 6, backgroundColor: '#dddddd', margin: 0 }} footer={null} destroyOnClose={true} closable={false} onCancel={actions.clearActive}>
          <video autoplay="true" controls src={hdUrl} width={state.width} height={state.height} 
              playsinline="true" />
        </Modal>
      </div>
    </VideoAssetWrapper>
  )
}

