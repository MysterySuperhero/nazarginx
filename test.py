#!/usr/bin/env python

import re
import socket
import httplib
import unittest

class HttpServer(unittest.TestCase):
  host = "localhost"
  port = 80

  def setUp(self):
    self.conn = httplib.HTTPConnection(self.host, self.port, timeout=10)

  def tearDown(self):
    self.conn.close()


  def test_document_root_escaping(self):
    """document root escaping forbidden"""
    self.conn.request("GET", "/httptest/../../../../../../../../../../../../../etc/passwd")
    r = self.conn.getresponse()
    data = r.read()
    self.assertIn(int(r.status), (400, 403, 404))


loader = unittest.TestLoader()
suite = unittest.TestSuite()
a = loader.loadTestsFromTestCase(HttpServer)
suite.addTest(a)

class NewResult(unittest.TextTestResult):
  def getDescription(self, test):
    doc_first_line = test.shortDescription()
    return doc_first_line or ""

class NewRunner(unittest.TextTestRunner):
  resultclass = NewResult

runner = NewRunner(verbosity=2)
runner.run(suite)

