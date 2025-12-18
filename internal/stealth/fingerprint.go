package stealth

import "github.com/go-rod/rod"

// MaskFingerprint modifies JS-exposed properties (demo-safe)
func MaskFingerprint(page *rod.Page) {
	page.MustEval(`
		() => {
			Object.defineProperty(navigator, 'webdriver', { get: () => undefined });
			Object.defineProperty(navigator, 'platform', { get: () => 'Win32' });
			Object.defineProperty(navigator, 'language', { get: () => 'en-US' });
			Object.defineProperty(navigator, 'languages', { get: () => ['en-US', 'en'] });

			console.log('[DEMO] Fingerprint masked');
		}
	`)
}
