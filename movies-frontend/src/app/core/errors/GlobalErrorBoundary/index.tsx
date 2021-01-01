import * as React from 'react';
import { ErrorInfo } from 'react';

export class GlobalErrorBoundary extends React.Component<any, any> {
  constructor(props: any) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error: Error) {
    return { hasError: true };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    // Log error to metrics
    console.log("Lo capture"+error)
  }

  render() {
    const { hasError, errorInfo } = this.state;
    if (hasError) {
      return (
        <div>
          <div>
            <p>
              An error has occurred in this component.{' '}
              <span
                style={{ cursor: 'pointer', color: '#0077FF' }}
                onClick={() => {
                  window.location.reload();
                }}
              >
                Reload this page
              </span>{' '}
            </p>
          </div>

          <div>
            <details>
              <summary>Click for error details</summary>
              {errorInfo && errorInfo.componentStack.toString()}
            </details>
          </div>
        </div>
      );
    }
    return this.props.children;

  }
}
