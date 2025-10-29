import XCTest
import SwiftTreeSitter
import TreeSitterDockerfile

final class TreeSitterDockerfileTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_dockerfile())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Dockerfile grammar")
    }
}
